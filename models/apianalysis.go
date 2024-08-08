package models

import (
	"time"

	"gorm.io/gorm"
)

type OpenAPI struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Version     string    `json:"version"`
	Format      string    `json:"format"`
	UserID      uint      `json:"user_id"`
	OpenAPI     string    `json:"openapi"`
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
