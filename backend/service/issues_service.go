package service

import (
	"pcb-backend/backend/models"
	"pcb-backend/backend/repository"
)

func CreateIssuesService(issue models.Issue) error {
	return repository.CreateIssuesRepo(issue)
}
func GetAllIssuesService() ([]models.Issue, error) {
	return repository.GetAllIssuesRepo()
}
func GetIssuesByIdService(id int) (models.Issue, error) {
	return repository.GetIssuesByIdRepo(id)
}
func UpdateIssueByIdService(id int, data models.Issue) error {
	return repository.UpdateIssueByIdRepo(id, data)
}
func DeleteIssueService(id int) error {
	return repository.DeleteIssueRepo(id)
}
