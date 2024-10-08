package models

import (
	"time"

	"github.com/lib/pq"
)

type Project struct {
	ID              uint             `json:"id" gorm:"primaryKey;autoIncrement"`
	Title           string           `json:"title"`
	Description     string           `json:"description"`
	Languages       pq.StringArray   `json:"languages" gorm:"type:text[]"`
	Backend         bool             `json:"backend"`
	Frontend        bool             `json:"frontend"`
	SnapshotCode    string           `json:"snapshot_code"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	UserID          uint             `json:"user_id"`
	User            User             `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // User has many projects
	Recommendations []Recommendation `json:"recommendations" gorm:"foreignKey:ProjectID"`                // One project has many recommendations
}
