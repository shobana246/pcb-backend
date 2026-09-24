package repository

import (
	"database/sql"
	"errors"
	"pcb-backend/backend/database"
	"pcb-backend/backend/models"

	"gorm.io/gorm"
)

func ComponentCreationRepo(component models.Component) error {
	return database.DB.Create(&component).Error
}

func GetComponentRepo(ProjectId int) ([]models.Component, error) {
	var components []models.Component

	err := database.DB.
		Where("project_id = ? AND is_deleted = ?", ProjectId, false).
		Find(&components).Error

	return components, err
}

func GetByIdComponentRepo(id int) (models.Component, error) {
	var data models.Component

	err := database.DB.
		Where("id = ? AND is_deleted = ?", id, false).
		First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, sql.ErrNoRows
	}
	return data, err
}

func UpdateComponentRepo(id int, component models.Component) error {
	result := database.DB.
		Model(&models.Component{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Select("ComponentName", "Package", "Category", "PlacementSide",
			"Rotation", "Xposition", "Yposition", "Height", "Supplier",
			"PartNumber", "Status", "TolerancePosition", "ToleranceRotation",
			"ImageUrl", "Notes").
		Updates(component)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func DeleteComponentRepo(id int) error {
	result := database.DB.
		Model(&models.Component{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Update("is_deleted", true)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func ProjectActiveRepo(id int) (bool, error) {
	var count int64
	err := database.DB.Table("projects").
		Where("id = ? AND is_deleted = ?", id, false).
		Count(&count).Error
	return count > 0, err
}

func ComponentNameExistsRepo(projectID int, name string, excludeID int) (bool, error) {
	var count int64
	err := database.DB.Table("component").
		Where("project_id = ? AND component_name = ? AND is_deleted = ? AND id <> ?", projectID, name, false, excludeID).
		Count(&count).Error
	return count > 0, err
}
