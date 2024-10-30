package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	OrderID   string `json:"order_id"`
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var orders = []Order{
	{OrderID: "101", UserID: "1", ProductID: "2"},
}

func getUser(userID string) (*User, error) {
	response, err := http.Get("http://user-service:8081/users/" + userID)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func getProduct(productID string) (*Product, error) {
	response, err := http.Get("http://product-service:8082/products/" + productID)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var product Product
	if err := json.NewDecoder(response.Body).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func main() {
	router := gin.Default()

	router.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, orders)
	})

	router.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, order := range orders {
			if order.OrderID == id {
				user, _ := getUser(order.UserID)
				product, _ := getProduct(order.ProductID)
				c.JSON(http.StatusOK, gin.H{
					"order":   order,
					"user":    user,
					"product": product,
				})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
	})

	router.Run(":8083")
}
