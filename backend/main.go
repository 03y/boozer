package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

const NAME      string = "🍺 boozer"
const VERSION   string = "0.1-Alpha"

func uptime(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "hello boozers!")
}

func main() {
    fmt.Printf("%s %s\n", NAME, VERSION)

    if len(os.Args) < 2 {
        fmt.Println("Usage: ./boozer <URL>:<PORT>")
        return
    }

    router := gin.Default()
    router.GET("/hello", uptime)

    var listen string = os.Args[1]
    fmt.Printf("\nLets get boozing! 🍻\nListening on %s...\n\n", listen)
    router.Run(listen)
}
