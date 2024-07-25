package models

import "time"

type Recommendation struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID   uint      `json:"project_id"` // Foreign key to relate to Project
	Project     Project   `json:"project" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Security    string    `json:"security"`    // Security recommendations
	Readability string    `json:"readability"` // Readability recommendations
	Other       string    `json:"other"`       // Other types of recommendations
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
