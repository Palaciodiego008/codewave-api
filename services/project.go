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

func ListProjects(userID uint) ([]models.Project, error) {
	projects := []models.Project{}
	if err := config.DB.Where("user_id = ?", userID).Preload("User").Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func UpdateProject(project *models.Project) error {
	return config.DB.Save(project).Error
}
