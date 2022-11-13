package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func HelloUser(c *gin.Context) {
  name := c.Param("name")
	c.String(http.StatusOK, "hello %s ", name)
}

func main() {
	log.Printf("Starting...")

	gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", Hello)
	router.GET("/user/:name", HelloUser)

	router.Run(":3100")
}