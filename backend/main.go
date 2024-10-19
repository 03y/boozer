package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"os"
	"strconv"
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

const NAME string = "🍺 boozer"
const VERSION string = "0.1-Alpha"

var db_connection *pgxpool.Conn

func connect_db() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	db_connection, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Printf("WARNING: Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to database!")
	defer db_connection.Close()
}

/* ******************************************************************************** */
/* API endpoints */
/* ******************************************************************************** */

func hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hello boozers!")
}

func get_item(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("item_id"))
	beer := &item{Item_id: id, Name: "Tennents", Units: 2.272, Added: 1729331086}

	serialised, err := json.Marshal(beer)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, "Internal error")
		return
	}
	c.IndentedJSON(http.StatusOK, string(serialised))
}

/* ******************************************************************************** */

func main() {
	fmt.Printf("%s %s\n", NAME, VERSION)

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./boozer <URL>:<PORT>")
		return
	}

	connect_db()

	// Setup API routes
	router := gin.Default()
	router.GET("/hello", hello)
	router.GET("/item/:item_id", get_item)

	var listen string = os.Args[1]
	fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)
	router.Run(listen)
}
