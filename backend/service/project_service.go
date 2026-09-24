package service

import (
	"pcb-backend/backend/models"
	"pcb-backend/backend/repository"
)

func ProjectCreationService(project_data models.Project) error {
	return repository.ProjectCreationRepo(project_data)
}

func GetAllServiceProject() ([]models.Project, error) {
	return repository.GetAllProjectRepo()
}

func GetProjectByIdService(id int) (models.Project, error) {
	return repository.GetProjectByIdRepo(id)
}

func ProjectUpdateService(id int, update_data models.Project) error {
	return repository.ProjectUpdateRepo(id, update_data)
}

func ProjectDeleteService(id int) error {
	return repository.ProjectDeleteRepo(id)
}
