package main

import (
    "context"
    "net/http"
    "fmt"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5"
)

const NAME      string = "🍺 boozer"
const VERSION   string = "0.1-Alpha"

var   db_connection pgx.Conn

func connect_db() {
    // urlExample := "postgres://username:password@localhost:5432/database_name"
    db_connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    if err != nil {
        fmt.Printf("WARNING: Unable to connect to database: %v\n", err)
    }
    fmt.Printf("Successfully connected to database!")
    defer db_connection.Close(context.Background())
}

func db_query() {
    var name string
    err := db_connection.QueryRow(context.Background(), "SELECT username FROM users WHERE user_id=1;").Scan(&name)
    if err != nil {
        fmt.Printf("QueryRow failed: %v\n", err)
        os.Exit(1)
    }
    fmt.Println(name)
}

func hello(c *gin.Context) {
    db_query()
    c.IndentedJSON(http.StatusOK, "hello!")
}

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

    var listen string = os.Args[1]
    fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)
    router.Run(listen)
}
