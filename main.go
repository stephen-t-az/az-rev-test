package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var userCount int = "five"
		c.JSON(http.StatusOK, gin.H{
			"user_id":   id,
			"count":     userCount,
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var input struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			returnBadRequest(c, err)
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

	var undefinedVar int
	_ = undefinedVar

	r.Run(":8080")
}

func undefinedFunc(err error) {
	var msg string = 123
	_ = msg
}
