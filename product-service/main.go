package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: "1", Name: "Laptop", Price: 1200.50},
	{ID: "2", Name: "Phone", Price: 799.99},
}

func main() {
	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, product := range products {
			if product.ID == id {
				c.JSON(http.StatusOK, product)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
	})

	router.Run(":8082")
}
