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

type TodoController struct {
  postgresPass string
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	db, err := gorm.Open(postgres.Open(c.postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var todos []Todo
  db.Find(&todos)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todos": todos,
	})
}

func (c *TodoController) ShowTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	db, err := gorm.Open(postgres.Open(c.postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var todo Todo
  db.First(&todo, id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": todo,
	})
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	db, err := gorm.Open(postgres.Open(c.postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var newTodo Todo
	ctx.BindJSON(&newTodo);
  db.Create(&newTodo)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": newTodo,
	})
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	db, err := gorm.Open(postgres.Open(c.postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var updateTodo Todo
  db.First(&updateTodo, id)
	ctx.BindJSON(&updateTodo);
	db.Save(&updateTodo)
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": updateTodo,
	})
} 

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	db, err := gorm.Open(postgres.Open(c.postgresPass), &gorm.Config{})
	if err != nil {
		errorMessage := err.Error()
		fmt.Println("DB接続エラー")
		panic(errorMessage)
	}
	fmt.Println("DB接続に成功しました。")

	var deleteTodo Todo
  db.Delete(&deleteTodo, id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": deleteTodo,
	})
}

func GetEnv() (string, string, string, string, error) {
	err := godotenv.Load("env/dev.env")
	if err != nil {
		return "", "", "", "", err
	}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	// host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DBNAME")
	
	return user, password, port, dbname, nil
}

func NewTodoController(postgresPass string) *TodoController {
  return &TodoController{
    postgresPass: postgresPass,
  }
}


func main() {
	router := gin.Default()
	USER, PASSWORD, PORT, DBNAME, err :=  GetEnv()
	if err != nil {
		fmt.Println("環境変数エラー", err)
		return
	}
	postgresPass := "postgresql://" + USER + ":" + PASSWORD + "@db" + ":" + PORT + "/" + DBNAME
	todoController := NewTodoController(postgresPass)

	router.GET("/todos", todoController.GetTodos)
	router.GET("/todo/:id", todoController.ShowTodo)
	router.POST("/todo", todoController.CreateTodo)
	router.PATCH("/todo/:id", todoController.UpdateTodo)
	router.DELETE("/todo/:id", todoController.DeleteTodo)

	router.Run(":8080")
}
