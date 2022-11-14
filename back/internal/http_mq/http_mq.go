package http_mq

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
}

func (h *HttpServer)Hello(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func (h *HttpServer)HelloUser(c *gin.Context) {
  name := c.Param("name")
	c.String(http.StatusOK, "hello %s ", name)
}

func (h *HttpServer) StartHttp(addr string) error {
	gin.DisableConsoleColor()
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", h.Hello)
	router.GET("/user/:name", h.HelloUser)
	err := router.Run(addr)
	return err
}