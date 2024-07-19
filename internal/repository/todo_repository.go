package repository

import (
    "log"
    "os"
    "path/filepath"
    "time"
    "todoapp/internal/model"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB
var err error

// InitDB initializes the database connection using SQLite and migrates the Todo model.
func InitDB() {
    dbPath := filepath.Join(os.TempDir(), "test.db") // AWS Lambdaの/tmpディレクトリを利用
    db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    log.Println("Database connected successfully.")

    err = db.AutoMigrate(&model.Todo{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }
    log.Println("Database migration completed.")
}

// GetAllTodos retrieves all todos from the database.
func GetAllTodos() []model.Todo {
    var todos []model.Todo
    result := db.Find(&todos)
    if result.Error != nil {
        log.Fatalf("Error fetching todos: %v", result.Error)
    }
    return todos
}

// CreateTodo inserts a new todo into the database.
func CreateTodo(todo *model.Todo) {
    result := db.Create(todo)
    if result.Error != nil {
        log.Fatalf("Error creating todo: %v", result.Error)
    }
    log.Println("Todo created successfully.")
}

// GetTodoByID retrieves a todo by its ID from the database.
// It returns the found todo and any error encountered.
func GetTodoByID(id string) (model.Todo, error) {
    var todo model.Todo
    result := db.First(&todo, "id = ?", id)
    if result.Error != nil {
        log.Printf("Error fetching todo by ID: %v", result.Error)
    }
    return todo, result.Error
}

// UpdateTodo updates an existing todo in the database.
// It returns any error encountered during the update.
func UpdateTodo(todo *model.Todo) error {
    result := db.Save(todo)
    if result.Error != nil {
        log.Printf("Error updating todo: %v", result.Error)
    }
    log.Println("Todo updated successfully.")
    return result.Error
}

// DeleteTodo removes a todo by its ID from the database.
// It returns any error encountered during the deletion.
func DeleteTodo(id string) error {
    result := db.Delete(&model.Todo{}, "id = ?", id)
    if result.Error != nil {
        log.Printf("Error deleting todo: %v", result.Error)
    }
    log.Println("Todo deleted successfully.")
    return result.Error
}

// SearchTodos searches for todos by title or description using a query string.
// It returns a slice of matching todos.
func SearchTodos(query string) []model.Todo {
    var todos []model.Todo
    result := db.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&todos)
    if result.Error != nil {
        log.Printf("Error searching todos: %v", result.Error)
    }
    return todos
}

// FinishTodoByID marks a todo as finished by its ID.
// It returns the updated todo and any error encountered.
func FinishTodoByID(id string) (model.Todo, error) {
    var todo model.Todo
    result := db.First(&todo, "id = ?", id)
    if result.Error != nil {
        log.Printf("Error fetching todo by ID: %v", result.Error)
        return todo, result.Error
    }

    now := time.Now()
    todo.FinishedAt = &now
    result = db.Save(&todo)
    if result.Error != nil {
        log.Printf("Error marking todo as finished: %v", result.Error)
    }
    log.Println("Todo marked as finished successfully.")
    return todo, result.Error
}
