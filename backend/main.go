package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"boozer/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/argon2"
)

type App struct {
	DB *pgx.Conn
}

type ArgonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

const NAME string = "🍺 boozer"
const VERSION string = "0.1-Alpha"

func hash(input string, params *ArgonParams) (encodedHash string, err error) {
	salt, err := generateRandomBytes(params.saltLength)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	hashed := argon2.IDKey([]byte(input), salt, params.iterations, params.memory, params.parallelism, params.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hashed)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, params.memory, params.iterations, params.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

/* ******************************************************************************** */
/* API endpoints */
/* ******************************************************************************** */

func (a *App) GetItem(c *gin.Context) {
	var beer models.Item
	err := a.DB.QueryRow(context.Background(), "SELECT * FROM items WHERE item_id=$1", c.Param("item_id")).Scan(&beer.Item_id, &beer.Name, &beer.Units, &beer.Added)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, beer)
}

func (a *App) GetItemList(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT * FROM items")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	beers := make([]models.Item, 0)
	for rows.Next() {
		var beer models.Item
		err := rows.Scan(&beer.Item_id, &beer.Name, &beer.Units, &beer.Added)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
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
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	if newBeer.Name == "" || newBeer.Units < 0 {
		c.JSON(http.StatusBadRequest, "Bad request data")
		return
	}

	newBeer.Added = int(time.Now().Unix())

	_, err = a.DB.Exec(context.Background(), "INSERT INTO items (name, units, time) VALUES ($1, $2, $3)", newBeer.Name, newBeer.Units, newBeer.Added)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	c.Status(http.StatusCreated)
}

func (a *App) AddUser(c *gin.Context) {
	var newUser models.User

	err := c.BindJSON(&newUser)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	if len(newUser.Username) < 1 || len(newUser.Username) > 20 {
		c.JSON(http.StatusBadRequest, "Bad request data")
		return
	}
	// todo: check decode hash

	newUser.Created = int(time.Now().Unix())

	_, err = a.DB.Exec(context.Background(), "INSERT INTO users (username, password, created) VALUES ($1, $2, $3)", newUser.Username, newUser.Password, newUser.Created)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	c.Status(http.StatusCreated)
}

func (a *App) GetUser(c *gin.Context) {
	var user models.User
	err := a.DB.QueryRow(context.Background(), "SELECT user_id, username, created FROM users WHERE user_id=$1", c.Param("user_id")).Scan(&user.User_id, &user.Username, &user.Created)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *App) AddConsumption(c *gin.Context) {
	var newConsumption models.Consumption
	err := c.BindJSON(&newConsumption)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	// write time here, dont allow user to mess with this
	// todo: in future maybe allow backdating
	newConsumption.Time = int(time.Now().Unix())

	_, err = a.DB.Exec(context.Background(), "INSERT INTO consumptions (user_id, item_id, time) VALUES ($1, $2, $3)", newConsumption.User_id, newConsumption.Item_id, newConsumption.Time)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	c.Status(http.StatusCreated)
}

func (a *App) GetConsumption(c *gin.Context) {
	var consumption models.Consumption
	err := a.DB.QueryRow(context.Background(), "SELECT consumption_id, item_id, user_id, time FROM consumptions WHERE consumption_id=$1", c.Param("consumption_id")).Scan(&consumption.Consumption_id, &consumption.Item_id, &consumption.User_id, &consumption.Time)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, consumption)
}

func (a *App) GetUserLeaderboard(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT users.username, COUNT(consumptions.item_id) AS drank FROM consumptions INNER JOIN users ON consumptions.user_id = users.user_id GROUP BY users.username ORDER BY drank DESC LIMIT 10;")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	leaderboard := make([]models.LeaderboardUser, 0)
	for rows.Next() {
		var user models.LeaderboardUser
		err := rows.Scan(&user.Username, &user.Consumed)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	leaderboard := make([]models.LeaderboardItem, 0)
	for rows.Next() {
		var item models.LeaderboardItem
		err := rows.Scan(&item.Item_id, &item.Name, &item.Units, &item.Added, &item.Consumed)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
			return
		}
		leaderboard = append(leaderboard, item)
	}

	c.JSON(http.StatusOK, leaderboard)
}

/* ******************************************************************************** */

func (a *App) setUpRouter() *gin.Engine {
	router := gin.Default()

	// adding new items and consumptions
	router.POST("/submit/item", a.AddItem)               // todo: maybe add field for who added it, add auth for this
	router.POST("/submit/consumption", a.AddConsumption) // todo: implement auth

	// getting items
	router.GET("/item/:item_id", a.GetItem)
	router.GET("/items", a.GetItemList)

	// account actions
	router.POST("/signup", a.AddUser)
	router.GET("/user/:user_id", a.GetUser)

	// get consumption
	//router.GET("/consumption/:consumption_id", a.GetConsumption) // todo: implement auth

	// leaderboards
	router.GET("/leaderboard/items", a.GetItemsLeaderboard)
	router.GET("/leaderboard/users", a.GetUserLeaderboard)

	return router
}

func main() {
	fmt.Printf("%s %s\n", NAME, VERSION)

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./boozer <URL>:<PORT>")
		return
	}

	// export DATABASE_URL='postgres://username:password@localhost:5432/database_name'
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("WARNING: Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to database!")
	defer db.Close(context.Background())
	app := &App{DB: db}

	router := app.setUpRouter()

	var listen string = os.Args[1]
	fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)
	router.Run(listen)
}
