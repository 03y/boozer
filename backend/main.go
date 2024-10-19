package main

import (
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
	When           int // unix timestamp
}

type App struct {
	DB *pgx.Conn
}

const NAME string = "🍺 boozer"
const VERSION string = "0.1-Alpha"

/* ******************************************************************************** */
/* API endpoints */
/* ******************************************************************************** */

func hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hello boozers!")
}

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
	router.GET("/hello", hello)
	router.GET("/item/:item_id", app.get_item)

	var listen string = os.Args[1]
	fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)
	router.Run(listen)
}
