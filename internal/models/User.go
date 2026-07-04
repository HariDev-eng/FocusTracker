package models

import "time"

// User is minimal for now — single-user app, no auth fields yet.
// UserID exists on Task so scoping ("give me *my* tasks") is already
// wired in before you ever add login.
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
