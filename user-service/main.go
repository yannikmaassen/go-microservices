package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: "1", Name: "Alice"},
	{ID: "2", Name: "Bob"},
}

func main() {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	})

	router.Run(":8081")
}
