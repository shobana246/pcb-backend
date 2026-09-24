package service

import (
	"pcb-backend/backend/models"
	"pcb-backend/backend/repository"
)

func ComponentCreationService(component models.Component) error {
	return repository.ComponentCreationRepo(component)
}
func GetComponentService(ProjectId int) ([]models.Component, error) {
	return repository.GetComponentRepo(ProjectId)
}
func GetByIdComponentService(Id int) (models.Component, error) {
	return repository.GetByIdComponentRepo(Id)
}
func UpdateComponentService(id int, Component models.Component) error {
	return repository.UpdateComponentRepo(id, Component)
}
func DeleteComponentService(id int) error {
	return repository.DeleteComponentRepo(id)
}
