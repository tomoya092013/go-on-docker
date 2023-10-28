package main

import (
	"fmt"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type Todo struct {
	Id uint
	Title string
	Description string
	Deadline string
	Status bool
	Priority int
}

func main() {
	router := gin.Default()
	router.GET("/todo/:id", func(c *gin.Context) {
		USER, PASSWORD, PORT, DBNAME :=  GetEnv()
		postgresPass := "postgresql://" + USER + ":" + PASSWORD + "@db" + ":" + PORT + "/" + DBNAME
		ShowTodo(c, postgresPass)
	})
	router.Run(":8080")
}

func ShowTodo(c *gin.Context, postgresPass string) {
	id := c.Param("id")
	
	db, err := gorm.Open(postgres.Open(postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}

	fmt.Println("DB接続に成功しました。")

	var todo Todo
  db.First(&todo, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": todo,
	})
}

func GetEnv() (string, string, string, string) {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		fmt.Println(".envファイルがありません")
	}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	// host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DBNAME")
	
	return user, password, port, dbname
}
