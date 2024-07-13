package services

import (
	"codewave/config"
	"codewave/models"
)

func CreateProject(project *models.Project) error {
	return config.DB.Create(project).Error
}

func GetProject(id string) (*models.Project, error) {
	var project models.Project
	if err := config.DB.Preload("User").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func ListProjects() ([]models.Project, error) {
	var projects []models.Project
	if err := config.DB.Preload("User").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
