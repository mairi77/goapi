package main

import (
	"log"
	"os"
	"todoapp/internal/controller"
	"todoapp/internal/repository"

	"github.com/akrylysov/algnhsa"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	log.Println("Initializing Database...")
	repository.InitDB()
	log.Println("Database Initialized.")

	router := gin.New() // DefaultをNewに変更してカスタムミドルウェアを追加
	router.Use(gin.Logger(), gin.Recovery())

	log.Println("Registering routes...")
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Todo App API!",
		})
	})
	router.GET("/todo", controller.GetTodos)
	router.POST("/todo", controller.CreateTodo)
	router.PUT("/todo/:id", controller.UpdateTodo)
	router.DELETE("/todo/:id", controller.DeleteTodo)
	router.PUT("/todo/:id/finish", controller.FinishTodo)
	router.GET("/todo/:id", controller.GetTodoByID)
	router.GET("/todo/search", controller.SearchTodos)
	log.Println("Routes registered.")

	return router
}

func main() {
	log.Println("Starting application...")

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		router := createRouter()
		log.Println("Running in lambda mode.")

		// 利用可能なフィールドをデフォルトのListenAndServeとして呼び出し
		algnhsa.ListenAndServe(router, &algnhsa.Options{
			UseProxyPath: true,
		})

	} else {
		router := createRouter()
		log.Println("Running in local mode.")
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}
}
