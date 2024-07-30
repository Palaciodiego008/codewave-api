package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Contact struct {
	Email string `json:"email"`
}

type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Info struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TermsOfService string  `json:"termsOfService"`
	Contact        Contact `json:"contact" gorm:"embedded"`
	License        License `json:"license" gorm:"embedded"`
	Version        string  `json:"version"`
}

type ExternalDocs struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Path struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type OpenAPI struct {
	gorm.Model
	ID        uint            `json:"id" gorm:"primaryKey"`
	UserID    uint            `json:"user_id"`
	OpenAPI   string          `json:"openapi"`
	Info      Info            `json:"info" gorm:"embedded"`
	Servers   pq.StringArray  `json:"servers" gorm:"type:text[]"`
	Tags      pq.StringArray  `json:"tags" gorm:"type:text[]"`
	Paths     map[string]Path `json:"paths" gorm:"type:jsonb"`
	User      User            `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
