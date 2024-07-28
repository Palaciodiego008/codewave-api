package models

import "time"

type Analysis struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProjectID uint      `json:"project_id"`
	Report    string    `json:"report"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GeminiResponse struct {
	Security           []Check `json:"Security,omitempty"`
	Readability        []Check `json:"Readability,omitempty"`
	StaticCodeAnalysis []Check `json:"Static Code Analysis,omitempty"`
	DependencyScanning []Check `json:"Dependency Scanning,omitempty"`
	SAST               []Check `json:"SAST,omitempty"`
}

type Check struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
