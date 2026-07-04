package models

import "time"

// Completion is one row per task per calendar day — this is the
// entire source of truth for streaks. No "last completed date" field
// on Task, because a single date can't tell you if yesterday was also
// done, or the day before that. You need the full history to walk
// backwards through and count.
//
// The composite unique index (same index name on both fields) stops
// you from ever accidentally creating two completion rows for the
// same task on the same day.

type Completion struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	TaskID            uint      `gorm:"index:idx_task_date,unique;not null" json:"task_id"`
	Date              time.Time `gorm:"index:idx_task_date,unique;type:date;not null" json:"date"`
	DistractionsCount int       `gorm:"default:0" json:"distractions_count"`
	Checkpoints       int       `gorm:"default:0" json:"checkpoints"`
	Note              string    `json:"note"`
	CreatedAt         time.Time `json:"created_at"`
}
