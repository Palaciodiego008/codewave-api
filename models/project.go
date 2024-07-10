package models

import "time"

type Project struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Language    string    `json:"language"`
	Backend     string    `json:"backend"`
	Frontend    string    `json:"frontend"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // User has many projects
}
