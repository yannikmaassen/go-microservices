package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users/*action", func(c *gin.Context) {
		action := c.Param("action")
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8081/users"+action)
	})

	router.GET("/products/*action", func(c *gin.Context) {
		action := c.Param("action")
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8082/products"+action)
	})

	router.GET("/orders/*action", func(c *gin.Context) {
		action := c.Param("action")
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8083/orders"+action)
	})

	router.Run(":8080")
}
