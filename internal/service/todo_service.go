// Package service implements the business logic for the todoapp.
package service

import (
	"time"
	"todoapp/internal/model"
	"todoapp/internal/repository"

	"github.com/google/uuid"
)

// GetAllTodos retrieves all todos using the repository layer.
func GetAllTodos() []model.Todo {
	return repository.GetAllTodos()
}

// CreateTodo creates a new todo with a unique ID and timestamps,
// and saves it using the repository layer.
func CreateTodo(todo *model.Todo) {
	todo.ID = uuid.NewString()  // Assign a new UUID as the ID.
	todo.CreatedAt = time.Now() // Set the current time as CreatedAt.
	repository.CreateTodo(todo) // Save the new todo.
}

// GetTodoByID retrieves a todo by its ID using the repository layer.
// It returns the found todo and any error encountered.
func GetTodoByID(id string) (model.Todo, error) {
	return repository.GetTodoByID(id)
}

// UpdateTodo updates an existing todo with timestamp
// and saves the updated todo using the repository layer.
// It returns any error encountered during the update.
func UpdateTodo(todo *model.Todo) error {
	todo.UpdatedAt = time.Now() // Set the current time as UpdatedAt.
	return repository.UpdateTodo(todo)
}

// DeleteTodo removes a todo by its ID using the repository layer.
// It returns any error encountered during the deletion.
func DeleteTodo(id string) error {
	return repository.DeleteTodo(id)
}

// FinishTodoByID marks a todo as finished by its ID using the repository layer.
// It returns the updated todo and any error encountered.
func FinishTodoByID(id string) (model.Todo, error) {
	return repository.FinishTodoByID(id)
}

// SearchTodos searches for todos by title or description using a query string
// and returns a slice of matching todos.
func SearchTodos(query string) []model.Todo {
	return repository.SearchTodos(query)
}
