package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type User struct {
	User_id  int
	Username string
	Joined   int // unix timestamp
}

type Item struct {
	Item_id int
	Name    string
	Units   float32
	Added   int // unix timestamp
}

type Consumption struct {
	Consumption_id int
	User_id        int
	Item_id        int
	Time           int // unix timestamp
}

type App struct {
	DB *pgx.Conn
}

const NAME string = "🍺 boozer"
const VERSION string = "0.1-Alpha"

/* ******************************************************************************** */
/* API endpoints */
/* ******************************************************************************** */

func (a *App) GetItem(c *gin.Context) {
	var beer Item
	err := a.DB.QueryRow(context.Background(), "SELECT * FROM Items WHERE Item_id=$1", c.Param("Item_id")).Scan(&beer.Item_id, &beer.Name, &beer.Units, &beer.Added)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	serialised, err := json.Marshal(beer)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
		return
	}
	c.JSON(http.StatusOK, string(serialised))
}

func (a *App) GetItemList(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT * FROM Items")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	beers := make([]Item, 0)
	for rows.Next() {
		var beer Item
		err := rows.Scan(&beer.Item_id, &beer.Name, &beer.Units, &beer.Added)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
			return
		}
		beers = append(beers, beer)
	}

	serialised, err := json.Marshal(beers)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
		return
	}
	c.JSON(http.StatusOK, string(serialised))
}

func (a *App) AddItem(c *gin.Context) {
	var newBeer Item
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

	_, err = a.DB.Exec(context.Background(), "INSERT INTO Items (name, units time) VALUES ($1, $2, $3)", newBeer.Name, newBeer.Units, newBeer.Added)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	c.Status(http.StatusCreated)
}

func (a *App) GetUser(c *gin.Context) {
	var User User
	err := a.DB.QueryRow(context.Background(), "SELECT User_id, Username, created FROM Users WHERE User_id=$1", c.Param("User_id")).Scan(&User.User_id, &User.Username, &User.Joined)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	serialised, err := json.Marshal(User)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
		return
	}
	c.JSON(http.StatusOK, string(serialised))
}

func (a *App) AddConsumption(c *gin.Context) {
	var newConsumption Consumption
	err := c.BindJSON(&newConsumption)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	// write time here, dont allow User to mess with this
	// todo: in future maybe allow backdating
	newConsumption.Time = int(time.Now().Unix())

	_, err = a.DB.Exec(context.Background(), "INSERT INTO Consumptions (User_id, Item_id, time) VALUES ($1, $2, $3)", newConsumption.User_id, newConsumption.Item_id, newConsumption.Time)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	c.Status(http.StatusCreated)
}

func (a *App) GetConsumption(c *gin.Context) {
	var Consumption Consumption
	err := a.DB.QueryRow(context.Background(), "SELECT Consumption_id, Item_id, User_id, time FROM Consumptions WHERE Consumption_id=$1", c.Param("Consumption_id")).Scan(&Consumption.Consumption_id, &Consumption.Item_id, &Consumption.User_id, &Consumption.Time)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	serialised, err := json.Marshal(Consumption)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
		return
	}
	c.JSON(http.StatusOK, string(serialised))
}

/* ******************************************************************************** */

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

	router := gin.Default()
	router.GET("/add/new", app.AddItem)
	router.GET("/item/:item_id", app.GetItem)
	router.GET("/items", app.GetItemList)
	router.GET("/user/:user_id", app.GetUser)
	router.GET("/consumption/:consumption_id", app.GetConsumption)
	router.POST("/consumption/new", app.AddConsumption)

	var listen string = os.Args[1]
	fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)
	router.Run(listen)
}
