package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"todoapp/internal/controller"
	"todoapp/internal/repository"

	"github.com/akrylysov/algnhsa"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	log.Println("Initializing Database...")
	repository.InitDB()
	log.Println("Database Initialized.")

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	log.Println("Registering routes...")
	router.GET("/", func(c *gin.Context) {
		log.Println("GET / endpoint hit!")
		c.JSON(200, gin.H{
			"message": "Welcome to the Todo App API!",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		log.Println("GET /ping endpoint hit!")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/todo", func(c *gin.Context) {
		log.Println("GET /todo endpoint hit!")
		controller.GetTodos(c)
	})
	router.POST("/todo", func(c *gin.Context) {
		log.Println("POST /todo endpoint hit!")
		controller.CreateTodo(c)
	})
	router.PUT("/todo/:id", func(c *gin.Context) {
		log.Println("PUT /todo/:id endpoint hit!")
		controller.UpdateTodo(c)
	})
	router.DELETE("/todo/:id", func(c *gin.Context) {
		log.Println("DELETE /todo/:id endpoint hit!")
		controller.DeleteTodo(c)
	})
	router.PUT("/todo/:id/finish", func(c *gin.Context) {
		log.Println("PUT /todo/:id/finish endpoint hit!")
		controller.FinishTodo(c)
	})
	router.GET("/todo/:id", func(c *gin.Context) {
		log.Println("GET /todo/:id endpoint hit!")
		controller.GetTodoByID(c)
	})
	router.GET("/todo/search", func(c *gin.Context) {
		log.Println("GET /todo/search endpoint hit!")
		controller.SearchTodos(c)
	})
	log.Println("Routes registered.")

	return router
}

func main() {
	log.Println("Starting application...")

	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		router := createRouter()
		log.Println("Running in lambda mode.")
		log.Printf("Lambda function name: %s", os.Getenv("AWS_LAMBDA_FUNCTION_NAME"))

		// 詳細なログ出力のためのカスタムミドルウェア
		router.Use(func(c *gin.Context) {
			log.Printf("Request event: Method=%s, URL=%s\n", c.Request.Method, c.Request.URL.String())
			c.Next()
		})

		algnhsa.ListenAndServe(router, &algnhsa.Options{
			UseProxyPath: true,
		})
	} else {
		log.Println("Running in local mode.")
		router := createRouter()
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}
}
