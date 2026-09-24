package repository

import (
	"database/sql"
	"pcb-backend/backend/database"
	"pcb-backend/backend/models"
)

func CreateIssuesRepo(issue models.Issue) error {
	return database.DB.Create(&issue).Error
}

func ComponentActiveRepo(id int) (bool, error) {
	var count int64
	err := database.DB.Table("component").
		Where("id = ? AND is_deleted = ?", id, false).
		Count(&count).Error
	return count > 0, err
}

func GetAllIssuesRepo() ([]models.Issue, error) {
	var issues []models.Issue
	err := database.DB.
		Table("placement_issues AS p").
		Select("p.id, p.component_id, c.component_name AS reference, p.issue, p.severity, p.status, p.created_at").
		Joins("JOIN component c ON p.component_id = c.id").
		Where("c.is_deleted = ?", false).
		Scan(&issues).Error
	return issues, err
}

func GetIssuesByIdRepo(id int) (models.Issue, error) {
	var data models.Issue
	result := database.DB.
		Table("placement_issues AS p").
		Select("p.id, p.component_id, c.component_name AS reference, p.issue, p.severity, p.status, p.created_at").
		Joins("JOIN component c ON p.component_id = c.id").
		Where("p.id = ? AND c.is_deleted = ?", id, false).
		Scan(&data)

	if result.Error != nil {
		return data, result.Error
	}
	if result.RowsAffected == 0 {
		return data, sql.ErrNoRows
	}
	return data, nil
}

func UpdateIssueByIdRepo(id int, status models.Status) error {
	result := database.DB.
		Model(&models.Issue{}).
		Where("id = ?", id).
		Update("status", status)

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
		Where("id = ?", id).
		Delete(&models.Issue{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
