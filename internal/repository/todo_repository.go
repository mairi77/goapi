// Package repository provides data access logic for the todoapp.
package repository

import (
	"time"
	"todoapp/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// InitDB initializes the database connection using SQLite and migrates the Todo model.
func InitDB() {
	dbPath := "/tmp/test.db" // AWS Lambdaの/tmpディレクトリを利用
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Todo{})
}

// GetAllTodos retrieves all todos from the database.
func GetAllTodos() []model.Todo {
	var todos []model.Todo
	db.Find(&todos)
	return todos
}

// CreateTodo inserts a new todo into the database.
func CreateTodo(todo *model.Todo) {
	db.Create(todo)
}

// GetTodoByID retrieves a todo by its ID from the database.
// It returns the found todo and any error encountered.
func GetTodoByID(id string) (model.Todo, error) {
	var todo model.Todo
	err := db.First(&todo, "id = ?", id).Error
	return todo, err
}

// UpdateTodo updates an existing todo in the database.
// It returns any error encountered during the update.
func UpdateTodo(todo *model.Todo) error {
	return db.Save(todo).Error
}

// DeleteTodo removes a todo by its ID from the database.
// It returns any error encountered during the deletion.
func DeleteTodo(id string) error {
	return db.Delete(&model.Todo{}, "id = ?", id).Error
}

// SearchTodos searches for todos by title or description using a query string.
// It returns a slice of matching todos.
func SearchTodos(query string) []model.Todo {
	var todos []model.Todo
	db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&todos)
	return todos
}

// FinishTodoByID marks a todo as finished by its ID.
// It returns the updated todo and any error encountered.
func FinishTodoByID(id string) (model.Todo, error) {
	var todo model.Todo
	if err := db.First(&todo, "id = ?", id).Error; err != nil {
		return todo, err
	}

	now := time.Now()
	todo.FinishedAt = &now
	db.Save(&todo)
	return todo, nil
}
