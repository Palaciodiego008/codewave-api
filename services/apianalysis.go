package services

import (
	"codewave/config"
	"codewave/models"
)

func CreateOpenAPI(openAPI *models.OpenAPI) error {
	err := config.DB.Create(openAPI).Error
	return err
}

func GetOpenAPIByID(id string) (*models.OpenAPI, error) {
	var openAPI models.OpenAPI

	err := config.DB.First(&openAPI, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &openAPI, nil
}
