package repository

import (
	"database/sql"
	"pcb-backend/backend/database"
	"pcb-backend/backend/models"
)

func CreateIssuesRepo(issue models.Issue) error {
	return database.DB.Create(&issue).Error

}

func GetAllIssuesRepo() ([]models.Issue, error) {

	var get_data []models.Issue
	err := database.DB.
		Table("placement_issues AS p").
		Select("p.id, c.component_name, p.issue, p.severity, p.status").
		Joins("JOIN component c ON p.component_id = c.id").
		Scan(&get_data).Error
	return get_data, err

}

func GetIssuesByIdRepo(id int) (data models.Issue, err error) {

	result := database.DB.
		Table("placement_issues as p").
		Select("p.id, c.component_name AS reference,p.issue, p.severity, p.status").
		Joins("JOIN component c ON p.component_id = c.id").
		Where("p.id = ?", id).
		Scan(&data)

	if result.Error != nil {
		return data, result.Error
	}
	if result.RowsAffected == 0 {
		return data, sql.ErrNoRows
	}

	return data, nil

}

func UpdateIssueByIdRepo(id int, data models.Issue) error {
	result := database.DB.
		Model(&models.Issue{}).
		Where("id = ?", id).
		Update("status", data.Status)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func DeleteIssueRepo(id int) error {

	result := database.DB.
		Where("id=?", id).
		Delete(&models.Issue{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
