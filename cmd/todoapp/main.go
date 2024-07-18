package main

import (
	"context"
	"os"
	"todoapp/internal/controller"
	"todoapp/internal/repository"

	"github.com/akrylysov/algnhsa" // GinフレームワークとLambdaを統合
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	repository.InitDB()

	router := gin.Default()

	router.GET("/todo", controller.GetTodos)
	router.POST("/todo", controller.CreateTodo)
	router.PUT("/todo/:id", controller.UpdateTodo)
	router.DELETE("/todo/:id", controller.DeleteTodo)
	router.PUT("/todo/:id/finish", controller.FinishTodo)
	router.GET("/todo/:id", controller.GetTodoByID)
	router.GET("/todo/search", controller.SearchTodos)

	return router
}

func lambdaHandler(ctx context.Context) (interface{}, error) {
	router := createRouter()
	algnhsa.ListenAndServe(router, nil)
	return nil, nil
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "lambda" {
		lambda.Start(lambdaHandler)
	} else {
		router := createRouter()
		router.Run(":8080")
	}
}
