package main

import (
	"os"
	"todoapp/internal/controller"
	"todoapp/internal/repository"

	"github.com/akrylysov/algnhsa" // GinフレームワークとLambdaを統合
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	repository.InitDB()

	// Ginのデフォルト設定を使用
	router := gin.Default()

	// ルートの定義
	router.GET("/todo", controller.GetTodos)
	router.POST("/todo", controller.CreateTodo)
	router.PUT("/todo/:id", controller.UpdateTodo)
	router.DELETE("/todo/:id", controller.DeleteTodo)
	router.PUT("/todo/:id/finish", controller.FinishTodo)
	router.GET("/todo/:id", controller.GetTodoByID)
	router.GET("/todo/search", controller.SearchTodos)

	return router
}

func main() {
	// コマンドライン引数に "lambda" が含まれていれば、Lambdaで実行されていると判断
	if len(os.Args) > 1 && os.Args[1] == "lambda" {
		router := createRouter()
		algnhsa.ListenAndServe(router, &algnhsa.Options{
			UseProxyPath: true,
		})
	} else {
		router := createRouter()
		router.Run(":8080") // ローカル環境で実行される場合
	}
}
