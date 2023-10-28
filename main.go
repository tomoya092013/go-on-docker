package main

import (
	"fmt"
	"net/http" 
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todo/:id", ShowTodo) 

	router.Run(":8080")
}

func ShowTodo(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id出力", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
