package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"index;not null" json:"user_id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Frequency   string    `gorm:"default:daily" json:"frequency"` // daily | weekdays | custom
	IsArchived  bool      `gorm:"default:false" json:"is_archived"`
	CreatedAt   time.Time `json:"created_at"`
}
