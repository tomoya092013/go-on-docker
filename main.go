package main

import (
	"fmt"
	"os"
	"net/http"
	// "reflect"

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
  db *gorm.DB
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	var todos []Todo
  if err := c.db.Find(&todos).Error; err != nil {
		c.handleError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todos": todos,
	})
}

func (c *TodoController) ShowTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var todo Todo
  if err := c.db.First(&todo, id).Error; err != nil {
		c.handleError(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": todo,
	})
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var newTodo Todo
	if err := ctx.BindJSON(&newTodo); err != nil {
		c.handleError(ctx, http.StatusInternalServerError, err)
	}
  if err := c.db.Create(&newTodo).Error; err != nil {
		c.handleError(ctx, http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": newTodo,
	})
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var updateTodo Todo
  if err := c.db.First(&updateTodo, id).Error; err != nil {
		c.handleError(ctx, http.StatusInternalServerError, err)
	}
	if err := ctx.BindJSON(&updateTodo); err != nil {
		c.handleError(ctx, http.StatusInternalServerError, err)
	}
	if err := c.db.Save(&updateTodo).Error;  err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todo": updateTodo,
	})
} 

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	var deleteTodo Todo
  if err := c.db.Delete(&deleteTodo, id).Error; err != nil {
		c.handleError(ctx, http.StatusInternalServerError, err)
	}

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

func (c *TodoController) handleError (ctx *gin.Context, status int, err error) {
	ctx.JSON(status, gin.H{
		"message": err.Error(),
	})
}

func NewTodoController(db *gorm.DB) *TodoController {
  return &TodoController{
    db: db,
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
	db, err := gorm.Open(postgres.Open(postgresPass), &gorm.Config{})
	if err != nil {
		fmt.Println("DB接続エラー", err)
		return
	}
	// defer db.Close()
	// fmt.Println("かた",reflect.TypeOf(db))

	todoController := NewTodoController(db)
	router.GET("/todos", todoController.GetTodos)
	router.GET("/todo/:id", todoController.ShowTodo)
	router.POST("/todo", todoController.CreateTodo)
	router.PATCH("/todo/:id", todoController.UpdateTodo)
	router.DELETE("/todo/:id", todoController.DeleteTodo)

	router.Run(":8080")
}
