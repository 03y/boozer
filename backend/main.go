package main

import (
	"context"
	"fmt"
	"os"
	"time"

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

type LeaderboardItem struct {
	TimesDrank int
	Item
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

	c.JSON(http.StatusOK, beers)
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

	_, err = a.DB.Exec(context.Background(), "INSERT INTO items (name, units time) VALUES ($1, $2, $3)", newBeer.Name, newBeer.Units, newBeer.Added)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Error processing data")
		return
	}
	c.Status(http.StatusCreated)
}

func (a *App) GetUser(c *gin.Context) {
	var user User
	err := a.DB.QueryRow(context.Background(), "SELECT user_id, username, created FROM users WHERE user_id=$1", c.Param("user_id")).Scan(&user.User_id, &user.Username, &user.Joined)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *App) AddConsumption(c *gin.Context) {
	var newConsumption Consumption
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
	var consumption Consumption
	err := a.DB.QueryRow(context.Background(), "SELECT consumption_id, item_id, user_id, time FROM consumptions WHERE consumption_id=$1", c.Param("consumption_id")).Scan(&consumption.Consumption_id, &consumption.Item_id, &consumption.User_id, &consumption.Time)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, consumption)
}

func (a *App) GetLeaderboard(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT items.item_id, items.name, items.units, items.added, COUNT(*) FROM consumptions INNER JOIN items ON consumptions.item_id = items.item_id GROUP BY items.item_id, items.name, items.units, items.added;")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	leaderboard := make([]LeaderboardItem, 0)
	for rows.Next() {
		var item LeaderboardItem
		err := rows.Scan(&item.Item_id, &item.Name, &item.Units, &item.Added, &item.TimesDrank)
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
	router.GET("/leaderboard", app.GetLeaderboard)

	var listen string = os.Args[1]
	fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)
	router.Run(listen)
}
