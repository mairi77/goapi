// Package controller handles HTTP requests for the todoapp.
package controller

import (
	"net/http"
	"todoapp/internal/model"
	"todoapp/internal/service"

	"github.com/gin-gonic/gin"
)

// GetTodos handles GET requests to retrieve all todos.
// It fetches all todos using the service layer and returns them in JSON format.
func GetTodos(c *gin.Context) {
	todos := service.GetAllTodos()
	c.JSON(http.StatusOK, todos)
}

// CreateTodo handles POST requests to create a new todo.
// It binds the JSON request body to a Todo model, creates the todo using the service layer,
// and returns the created todo in JSON format.
func CreateTodo(c *gin.Context) {
	var newTodo model.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.CreateTodo(&newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// UpdateTodo handles PUT requests to update an existing todo.
// It binds the JSON request body to a Todo model, updates the todo using the service layer,
// and returns the updated todo in JSON format.
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = id

	if err := service.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo handles DELETE requests to delete a todo by its ID.
// If the todo is not found, it returns a 404 error.
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteTodo(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// FinishTodo handles PUT requests to mark a todo as finished.
// If the todo is not found, it returns a 404 error.
func FinishTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := service.FinishTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// GetTodoByID handles GET requests to retrieve a todo by its ID.
// If the todo is not found, it returns a 404 error.
func GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := service.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// SearchTodos handles GET requests to search todos by a query.
// It fetches the matching todos using the service layer and returns them in JSON format.
func SearchTodos(c *gin.Context) {
	query := c.Query("q")
	todos := service.SearchTodos(query)
	c.JSON(http.StatusOK, todos)
}
