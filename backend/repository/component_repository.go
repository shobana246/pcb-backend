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
		Select("id,project_id, component_name, package, placement_side, rotation, x_position, y_position, height, supplier, part_number, status, tolerance_position, tolerance_rotation, notes").
		Where("project_id = ? AND is_deleted = ?", ProjectId, false).
		Find(&components).Error

	return components, err
}

func GetByIdComponentRepo(id int) (models.Component, error) {
	var data models.Component

	err := database.DB.
		Select("id,project_id, component_name, package, placement_side, rotation, x_position, y_position, height, supplier, part_number, status, tolerance_position, tolerance_rotation, notes").
		Where("id = ? AND is_deleted = ?", id, false).
		First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, sql.ErrNoRows
	}

	return data, nil
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
		Where("id = ?", id).
		Update("is_deleted", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
