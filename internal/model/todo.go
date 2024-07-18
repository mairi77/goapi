// Package model defines the data structures for the todoapp.
package model

import "time"

// Todo represents a to-do task.
type Todo struct {
	ID          string     `json:"id" gorm:"type:uuid;primary_key"`  // ID is the unique identifier for the todo, generated as a UUID.
	Title       string     `json:"title" binding:"required,max=128"` // Title is the name of the todo, required and cannot exceed 128 characters.
	Description string     `json:"description" binding:"required"`   // Description provides detailed information about the todo, required field.
	CreatedAt   time.Time  `json:"createdAt"`                        // CreatedAt is the timestamp when the todo was created.
	UpdatedAt   time.Time  `json:"updatedAt"`                        // UpdatedAt is the timestamp when the todo was last updated.
	FinishedAt  *time.Time `json:"finishedAt" sql:"index"`           // FinishedAt is the timestamp when the todo was completed, can be null.
}
