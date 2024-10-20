package main

import (
	"time"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
)

type user struct {
	User_id  int
	Username string
	Joined   int // unix timestamp
}

type item struct {
	Item_id int
	Name    string
	Units   float32
	Added   int // unix timestamp
}

type consumption struct {
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

func (a *App) get_item(c *gin.Context) {
	var beer item
	err := a.DB.QueryRow(context.Background(), "SELECT * FROM items WHERE item_id=$1", c.Param("item_id")).Scan(&beer.Item_id, &beer.Name, &beer.Units, &beer.Added)
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

func (a *App) get_item_list(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT * FROM items")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	beers := make([]item, 0)
	for rows.Next() {
		var beer item
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


func (a *App) get_user(c *gin.Context) {
	var user user
	err := a.DB.QueryRow(context.Background(), "SELECT user_id, username, created FROM users WHERE user_id=$1", c.Param("user_id")).Scan(&user.User_id, &user.Username, &user.Joined)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Error fetching data"})
		return
	}

	serialised, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing data"})
		return
	}
	c.JSON(http.StatusOK, string(serialised))
}

func (a *App) add_consumption(c *gin.Context) {
	var newConsumption consumption
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

/* ******************************************************************************** */

func main() {
	fmt.Printf("%s %s\n", NAME, VERSION)

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./boozer <URL>:<PORT>")
		return
	}

	// postgres://username:password@localhost:5432/database_name
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("WARNING: Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to database!")
	defer db.Close(context.Background())
	app := &App{DB: db}

	router := gin.Default()
	router.GET("/item/:item_id", app.get_item)
	router.GET("/items/", app.get_item_list)
	router.GET("/user/:user_id", app.get_user)
	router.POST("/consumption/new", app.add_consumption)

	var listen string = os.Args[1]
	fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)
	router.Run(listen)
}
