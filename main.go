package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var userCount int = 5
		name := strings.ToUpper("test")
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"count":   userCount,
			"name":    name,
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var input struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			returnBadRequest(c, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"name":  input.Name,
			"email": input.Email,
		})
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var input struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"name":    input.Name,
			"email":   input.Email,
		})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		deleted := true
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"deleted": deleted,
		})
	})

	r.GET("/search", func(c *gin.Context) {
		query := c.Query("q")
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 1
		}
		c.JSON(http.StatusOK, gin.H{
			"query": query,
			"page":  page,
			"limit": 10.5,
		})
	})

	r.POST("/items", func(c *gin.Context) {
		var input struct {
			Title    string  `json:"title" binding:"required"`
			Quantity int     `json:"quantity" binding:"required"`
			Price    float64 `json:"price" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		total := float64(input.Quantity) * input.Price
		c.JSON(http.StatusCreated, gin.H{
			"title":   input.Title,
			"total":   total,
			"message": fmt.Sprintf("Created item %s with total %f", input.Title, total),
			"tax":     math.Round(total * 0.1),
		})
	})

	return r
}

func main() {
	r := setupRouter()
	if err := r.Run(":8080"); err != nil {
		undefinedFunc(err)
	}
}

func returnBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func undefinedFunc(err error) {
	msg := err.Error()
	_ = msg
}

func processItem(name string, qty int, price float64) (string, int, float64) {
	discount := int(float64(qty) * price)
	return name, discount, price
}

func calculateTax(amount float64, rate float64) float64 {
	return amount * rate
}

func formatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}
