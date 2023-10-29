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
	USER, PASSWORD, PORT, DBNAME :=  GetEnv()
	postgresPass := "postgresql://" + USER + ":" + PASSWORD + "@db" + ":" + PORT + "/" + DBNAME

	router.GET("/todos", func(c *gin.Context) {
		GetTodos(c, postgresPass)
	})

	router.GET("/todo/:id", func(c *gin.Context) {
		ShowTodo(c, postgresPass)
	})

	router.POST("/todo", func(c *gin.Context) {
		CreateTodo(c, postgresPass)
	})

	router.PATCH("/todo/:id", func(c *gin.Context) {
		UpdateTodo(c, postgresPass)
	})

	router.DELETE("/todo/:id", func(c *gin.Context) {
		DeleteTodo(c, postgresPass)
	})

	router.Run(":8080")
}

func GetTodos(c *gin.Context, postgresPass string) {
	db, err := gorm.Open(postgres.Open(postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var todos []Todo
  db.Find(&todos)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todos": todos,
	})
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

func CreateTodo(c *gin.Context, postgresPass string) {
	db, err := gorm.Open(postgres.Open(postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var newTodo Todo
	c.BindJSON(&newTodo);
  db.Create(&newTodo)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": newTodo,
	})
}

func UpdateTodo(c *gin.Context, postgresPass string) {
	id := c.Param("id")
	db, err := gorm.Open(postgres.Open(postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var updateTodo Todo
  db.First(&updateTodo, id)
	c.BindJSON(&updateTodo);
	db.Save(&updateTodo)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": updateTodo,
	})
} 

func DeleteTodo(c *gin.Context, postgresPass string) {
	id := c.Param("id")
	db, err := gorm.Open(postgres.Open(postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var deleteTodo Todo
  db.Delete(&deleteTodo, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": deleteTodo,
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
