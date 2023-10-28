package main

import (
	"fmt"
	"net/http" 
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.GET("/todo/:id", ShowTodo)
	PrintEnv()

	router.Run(":8080")
}

func ShowTodo(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("id出力", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func PrintEnv() {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		fmt.Println(".envファイルがありません")
	}
	postgress_db := os.Getenv("POSTGRES_DB")
	postgress_user := os.Getenv("POSTGRES_USER")

	fmt.Println("DB", postgress_db)
	fmt.Println("ユーザー", postgress_user)
}
