package repository

import (
	"database/sql"
	"pcb-backend/backend/database"
	"pcb-backend/backend/models"

	"gorm.io/gorm"
)

func ProjectCreationRepo(creation_data models.Project) error {
	return database.DB.Create(&creation_data).Error
}

func GetAllProjectRepo() (get_all_data []models.Project, err error) {
	err = database.DB.
		Table("projects AS p").
		Select("p.id, p.project_name, p.revision, p.updated_at, p.updated_by, u.name AS updated_by_name").
		Joins("JOIN users u ON p.updated_by=u.id").
		Where("p.is_deleted = ?", false).
		Scan(&get_all_data).Error

	return get_all_data, err
}

func GetProjectByIdRepo(id int) (data models.Project, err error) {
	result := database.DB.
		Table("projects AS p").
		Select("p.id, p.project_name, p.revision, p.updated_at, p.updated_by, u.name AS updated_by_name").
		Joins("JOIN users u ON p.updated_by=u.id").
		Where("p.id = ? AND p.is_deleted = ?", id, false).
		Scan(&data)

	if result.Error != nil {
		return data, result.Error
	}
	if result.RowsAffected == 0 {
		return data, sql.ErrNoRows
	}
	return data, nil
}

func ProjectUpdateRepo(id int, update_data models.Project) error {
	result := database.DB.
		Model(&models.Project{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Select("project_name", "revision", "updated_by").
		Updates(&update_data)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func ProjectDeleteRepo(id int) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&models.Project{}).
			Where("id = ? AND is_deleted = ?", id, false).
			Update("is_deleted", true)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return sql.ErrNoRows
		}

		return tx.Model(&models.Component{}).
			Where("project_id = ?", id).
			Update("is_deleted", true).Error
	})
}

func UserExistsRepo(id int) (bool, error) {
	var count int64
	err := database.DB.Table("users").Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

func ProjectExistsRepo(name, revision string, excludeID int) (bool, error) {
	var count int64
	err := database.DB.Table("projects").
		Where("project_name = ? AND revision = ? AND is_deleted = ? AND id <> ?", name, revision, false, excludeID).
		Count(&count).Error
	return count > 0, err
}
