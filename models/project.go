package models

import "time"

type Project struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Backend     bool      `json:"backend"`
	Frontend    bool      `json:"frontend"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // User has many projects
}
