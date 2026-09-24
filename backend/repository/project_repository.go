package repository

import (
	"database/sql"
	"pcb-backend/backend/database"
	"pcb-backend/backend/models"
)

func ProjectCreationRepo(creation_data models.Project) error {
	return database.DB.Create(&creation_data).Error

}

func GetAllProjectRepo() (get_all_data []models.Project, err error) {

	err = database.DB.
		Table("projects AS p").
		Select("p.id, p.project_name, p.revision, p.updated_at, u.name AS updatedby").
		Joins("JOIN users u ON p.updated_by=u.id").
		Scan(&get_all_data).Error

	return get_all_data, err
}

func GetProjectByIdRepo(id int) (data models.Project, err error) {

	err = database.DB.
		Table("projects AS p").
		Select("p.id, p.project_name, p.revision, p.updated_at, u.name AS updatedby").
		Joins("JOIN users u On p.updated_by=u.id").
		Where("p.id=?", id).
		Scan(&data).Error

	return data, err
}

func ProjectUpdateRepo(id int, update_data models.Project) error {

	result := database.DB.
		Model(&models.Project{}).
		Where("id =?", id).
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
	err := database.DB.
		Model(&models.Component{}).
		Where("project_id = ?", id).
		Update("is_deleted", true).Error
	if err != nil {
		return err
	}

	result := database.DB.
		Model(&models.Project{}).
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
