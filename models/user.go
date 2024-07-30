package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Projects  []Project `json:"projects" gorm:"foreignKey:UserID"`
	OpenAPIs  []OpenAPI `json:"openapis" gorm:"foreignKey:UserID"`
}
