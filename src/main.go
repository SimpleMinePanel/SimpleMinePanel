package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var port = flag.Int("port", 8080, "specify the port the server will listen on")
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"value": "pong",
		})
	})

	flag.Parse()
	log.Default().Println("SMP is starting up...")
	log.Default().Printf("Port is :%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
