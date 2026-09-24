package service

import (
	"database/sql"
	"errors"
	"pcb-backend/backend/models"
	"pcb-backend/backend/repository"
	"strings"
)

var (
	ErrInvalidInput    = errors.New("invalid input")
	ErrUserNotFound    = errors.New("user not found")
	ErrProjectExists   = errors.New("project already exists")
	ErrProjectNotFound = errors.New("project not found")
)

func checkProject(p *models.Project, excludeID int) error {

	p.ProjectName = strings.TrimSpace(p.ProjectName)
	p.Revision = strings.TrimSpace(p.Revision)
	if p.ProjectName == "" || p.Revision == "" {
		return ErrInvalidInput
	}

	userOk, err := repository.UserExistsRepo(p.UpdatedBy)
	if err != nil {
		return err
	}
	if !userOk {
		return ErrUserNotFound
	}

	exists, err := repository.ProjectExistsRepo(p.ProjectName, p.Revision, excludeID)
	if err != nil {
		return err
	}
	if exists {
		return ErrProjectExists
	}
	return nil
}

func ProjectCreationService(project models.Project) error {
	if err := checkProject(&project, 0); err != nil {
		return err
	}
	return repository.ProjectCreationRepo(project)
}

func GetAllServiceProject() ([]models.Project, error) {
	return repository.GetAllProjectRepo()
}

func GetProjectByIdService(id int) (models.Project, error) {
	project, err := repository.GetProjectByIdRepo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return project, ErrProjectNotFound
	}
	return project, err
}

func ProjectUpdateService(id int, project models.Project) error {
	if err := checkProject(&project, id); err != nil {
		return err
	}
	err := repository.ProjectUpdateRepo(id, project)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrProjectNotFound
	}
	return err
}

func ProjectDeleteService(id int) error {
	err := repository.ProjectDeleteRepo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrProjectNotFound
	}
	return err
}
