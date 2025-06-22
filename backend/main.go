package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/subtle"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"boozer/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type App struct {
	DB      *pgxpool.Pool
	JWT_KEY *ecdsa.PrivateKey
}

const NAME string = "🍺 boozer"
const VERSION string = "0.3-Alpha"

var ARGON2_PARAMS = &params{
	memory:      65536,
	iterations:  3,
	parallelism: 2,
	saltLength:  128,
	keyLength:   128,
}

func plaintextToEncodedHash(input string) (encodedHash string, err error) {
	salt, err := generateRandomBytes(ARGON2_PARAMS.saltLength)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	hashed := argon2.IDKey([]byte(input), salt, ARGON2_PARAMS.iterations, ARGON2_PARAMS.memory, ARGON2_PARAMS.parallelism, ARGON2_PARAMS.keyLength)

	b64Salt := base64.StdEncoding.EncodeToString(salt)
	b64Hash := base64.StdEncoding.EncodeToString(hashed)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, ARGON2_PARAMS.memory, ARGON2_PARAMS.iterations, ARGON2_PARAMS.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func hash(input string) (hash []byte, err error) {
	salt, err := generateRandomBytes(ARGON2_PARAMS.saltLength)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return argon2.IDKey([]byte(input), salt, ARGON2_PARAMS.iterations, ARGON2_PARAMS.memory, ARGON2_PARAMS.parallelism, ARGON2_PARAMS.keyLength), nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func comparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.StdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.StdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func loadKey(keyFile string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing the private key")
	}

	return x509.ParseECPrivateKey(block.Bytes)
}

func (a *App) generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour).Unix(), // expire in 1 hour
	})

	return token.SignedString(a.JWT_KEY)
}

func parseJWT(tokenString string, privateKey *ecdsa.PrivateKey) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return &privateKey.PublicKey, nil
	})

	switch {
	case token.Valid:
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("Failed to parse token claims")
		} else {
			return claims, nil
		}
	default:
		return nil, err
	}
}

/* ******************************************************************************** */
/* API endpoints */
/* ******************************************************************************** */

func (a *App) GetItem(c *gin.Context) {
	var beer models.Item
	err := a.DB.QueryRow(context.Background(), "SELECT * FROM items WHERE item_id=$1", c.Param("item_id")).Scan(&beer.Item_id, &beer.Name, &beer.Units, &beer.Added)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println(err)
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, beer)
}

func (a *App) GetItemList(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT * FROM items ORDER BY added DESC")
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	beers := make([]models.Item, 0)
	for rows.Next() {
		var beer models.Item
		err := rows.Scan(&beer.Item_id, &beer.Name, &beer.Units, &beer.Added)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		beers = append(beers, beer)
	}

	c.JSON(http.StatusOK, beers)
}

func (a *App) AddItem(c *gin.Context) {
	var newBeer models.Item
	err := c.BindJSON(&newBeer)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	if newBeer.Name == "" || newBeer.Units < 0 || len(newBeer.Name) < 1 || len(newBeer.Name) > 40 {
		c.Status(http.StatusBadRequest)
		return
	}

	newBeer.Added = int(time.Now().Unix())

	_, err = a.DB.Exec(context.Background(), "INSERT INTO items (name, units, added) VALUES ($1, $2, $3)", newBeer.Name, newBeer.Units, newBeer.Added)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest) // it was probably the clients fault
		return
	}

	c.Status(http.StatusCreated)
}

func (a *App) AddUser(c *gin.Context) {
	var newUser models.User

	err := c.BindJSON(&newUser)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if len(newUser.Username) < 1 || len(newUser.Username) > 20 {
		c.Status(http.StatusBadRequest)
		return
	}

	hashedPassword, err := plaintextToEncodedHash(newUser.Password)

	newUser.Created = int(time.Now().Unix())

	_, err = a.DB.Exec(context.Background(), "INSERT INTO users (username, password, created) VALUES ($1, $2, $3)", newUser.Username, hashedPassword, newUser.Created)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (a *App) Authenticate(c *gin.Context) {
	// get user data from req
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// get the hash we have in the db
	var storedHash string
	err = a.DB.QueryRow(context.Background(), "SELECT password FROM users WHERE username=$1", user.Username).Scan(&storedHash)

	// compare
	match, err := comparePasswordAndHash(user.Password, storedHash)
	if match {
		// create jwt
		token, err := a.generateJWT(user)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		// TODO: we should tell the user (moreso the frontend application) how long their token is valid for (when exp is implemented)
		c.JSON(http.StatusOK, gin.H{"token": token})
		fmt.Println("Successful auth for user", user.Username)
	} else {
		c.Status(http.StatusBadRequest)
		return
	}
}

func (a *App) GetUser(c *gin.Context) {
	var user models.UserNoPw
	err := a.DB.QueryRow(context.Background(), "SELECT user_id, username, created FROM users WHERE user_id=$1", c.Param("user_id")).Scan(&user.User_id, &user.Username, &user.Created)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println(err)
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *App) GetUsername(c *gin.Context) {
	tokenString := c.Request.Header["Authorization"][0]
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1) // TODO: do this in other places too

	claims, err := parseJWT(tokenString, a.JWT_KEY)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, claims)
}

func (a *App) AddConsumption(c *gin.Context) {
	tokenString := c.Request.Header["Authorization"][0]
	claims, err := parseJWT(tokenString, a.JWT_KEY)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	var newConsumption models.Consumption
	err = c.BindJSON(&newConsumption)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// check item exists
	var itemId int
	err = a.DB.QueryRow(context.Background(), "SELECT item_id FROM items WHERE item_id=$1", newConsumption.Item_id).Scan(&itemId)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println(err)
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	// if no rows are returned then it doesnt exist
	// otherwise we are safe to continue knowing the item id exists

	// write time here, dont allow user to mess with this
	// TODO: in future maybe allow backdating
	newConsumption.Time = int(time.Now().Unix())

	var id_lookup string
	err = a.DB.QueryRow(context.Background(), "SELECT user_id FROM users WHERE username=$1", claims["username"]).Scan(&id_lookup)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	newConsumption.User_id, _ = strconv.Atoi(id_lookup)

	_, err = a.DB.Exec(context.Background(), "INSERT INTO consumptions (user_id, item_id, time) VALUES ($1, $2, $3)", newConsumption.User_id, newConsumption.Item_id, newConsumption.Time)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (a *App) RemoveConsumption(c *gin.Context) {
	// do auth
	tokenString := c.Request.Header["Authorization"][0]
	claims, err := parseJWT(tokenString, a.JWT_KEY)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// get the consumption requested for deletion (although we will only read the id)
	var newConsumption models.Consumption
	err = c.BindJSON(&newConsumption)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// get username from consumption id, check that is the authenticated user
	var usernameLookup string
	err = a.DB.QueryRow(context.Background(), "SELECT users.username FROM consumptions INNER JOIN users ON consumptions.user_id=users.user_id WHERE consumptions.consumption_id=$1", newConsumption.Consumption_id).Scan(&usernameLookup)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// check authenticated user is the user associated with consumption
	if usernameLookup != claims["username"] {
		fmt.Println("User", claims["username"], "tried to delete a consumption record from", usernameLookup, "!")
		c.Status(http.StatusBadRequest)
		return
	}

	_, err = a.DB.Exec(context.Background(), "DELETE FROM consumptions WHERE consumption_id=$1", newConsumption.Consumption_id)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (a *App) GetConsumption(c *gin.Context) {
	// auth first
	tokenString := c.Request.Header["Authorization"][0]
	claims, err := parseJWT(tokenString, a.JWT_KEY)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	// now get that consumption (presumably it exists!)
	var consumption models.Consumption
	var usernameLookup string
	err = a.DB.QueryRow(context.Background(), "SELECT consumptions.consumption_id, consumptions.item_id, consumptions.user_id, users.username, consumptions.time FROM consumptions INNER JOIN users ON consumptions.user_id=users.user_id WHERE consumptions.consumption_id=$1", c.Param("consumption_id")).Scan(&consumption.Consumption_id, &consumption.Item_id, &consumption.User_id, &usernameLookup, &consumption.Time)
	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Println(err)
		}
		c.Status(http.StatusNotFound)
		return
	}

	// check user matches
	if claims["username"] != usernameLookup {
		fmt.Println("authenticated user didnt match consumptions user")
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, consumption)
}

func (a *App) GetUserLeaderboard(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT users.username, COUNT(consumptions.item_id) AS drank FROM consumptions INNER JOIN users ON consumptions.user_id = users.user_id GROUP BY users.username ORDER BY drank DESC LIMIT 10;")
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusNotFound)
		return
	}

	leaderboard := make([]models.LeaderboardUser, 0)
	for rows.Next() {
		var user models.LeaderboardUser
		err := rows.Scan(&user.Username, &user.Consumed)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusNotFound)
			return
		}
		leaderboard = append(leaderboard, user)
	}

	c.JSON(http.StatusOK, leaderboard)
}

func (a *App) GetItemsLeaderboard(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT items.item_id, items.name, items.units, items.added, COUNT(items.item_id) AS drank FROM consumptions INNER JOIN items ON consumptions.item_id = items.item_id GROUP BY items.item_id, items.name, items.units, items.added ORDER BY drank DESC LIMIT 50;")
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusNotFound)
		return
	}

	leaderboard := make([]models.LeaderboardItem, 0)
	for rows.Next() {
		var item models.LeaderboardItem
		err := rows.Scan(&item.Item_id, &item.Name, &item.Units, &item.Added, &item.Consumed)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusNotFound)
			return
		}
		leaderboard = append(leaderboard, item)
	}

	c.JSON(http.StatusOK, leaderboard)
}

/* ******************************************************************************** */

func (a *App) setUpRouter() *gin.Engine {
	router := gin.Default()

	// cors
	config := cors.Config{
		AllowOrigins:     []string{"https://192.168.0.34:6969", "https://192.168.0.34"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	// adding new items/consumptions
	router.POST("/submit/item", a.AddItem) // TODO: maybe add field for who added it, add auth for this
	router.POST("/submit/consumption", a.AddConsumption)
	// TODO: router.PUT("/submit/consumption", a.AddConsumption)

	// updating and deleting items/consumptions
	router.POST("/remove/consumption", a.RemoveConsumption)

	// getting items
	router.GET("/item/:item_id", a.GetItem)
	router.GET("/items", a.GetItemList)

	// create & authenticate accounts
	router.POST("/signup", a.AddUser)
	router.POST("/authenticate", a.Authenticate)

	// get user
	router.GET("/user/:user_id", a.GetUser)
	router.GET("/username", a.GetUsername)

	// get consumption
	router.GET("/consumption/:consumption_id", a.GetConsumption) // TODO: implement auth

	// leaderboards
	router.GET("/leaderboard/items", a.GetItemsLeaderboard)
	router.GET("/leaderboard/users", a.GetUserLeaderboard)

	return router
}

func main() {
	fmt.Printf("%s %s\n", NAME, VERSION)

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./boozer <URL>:<PORT>")
		fmt.Println("Note: cert.pem and key.pem must exist in the current directory")
		return
	}

	// Check if SSL certificate files exist
	if _, err := os.Stat("cert.pem"); os.IsNotExist(err) {
		fmt.Println("cert.pem not found")
	}
	if _, err := os.Stat("key.pem"); os.IsNotExist(err) {
		fmt.Println("key.pem not found")
	}

	jwtKey, err := loadKey("boozer.pem")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// DATABASE_URL='postgres://username:password@localhost:5432/database_name'
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("WARNING: Error parsing config:", err)
		os.Exit(1)
	}

	// pool config
	config.MaxConns = 25
	config.MinConns = 25
	// TODO: find the best values for these
	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	// connect
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Failed to create connection pool:", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to database!")
	defer pool.Close()

	app := &App{DB: pool, JWT_KEY: jwtKey}

	router := app.setUpRouter()

	var listen string = os.Args[1]
	fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)

	srv := &http.Server{
		Addr:    listen,
		Handler: router,
	}

	fmt.Println("Starting HTTPS server on", listen)
	if err := srv.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && err != http.ErrServerClosed {
		fmt.Println("Error starting HTTPS server:", err)
	}
}
