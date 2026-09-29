package service

import (
	"database/sql"
	"errors"
	"pcb-backend/backend/models"
	"pcb-backend/backend/repository"
	"strings"
)

var (
	ErrComponentExists   = errors.New("component already exists in this project")
	ErrComponentNotFound = errors.New("component not found")
)

func ComponentCreationService(component models.Component) error {
	if err := cleanComponent(&component); err != nil {
		return err
	}
	if component.Status == nil {
		s := models.Unverified
		component.Status = &s
	}

	active, err := repository.ProjectActiveRepo(component.ProjectId)
	if err != nil {
		return err
	}
	if !active {
		return ErrProjectNotFound
	}

	exists, err := repository.ComponentNameExistsRepo(component.ProjectId, component.ComponentName, 0)
	if err != nil {
		return err
	}
	if exists {
		return ErrComponentExists
	}

	return repository.ComponentCreationRepo(component)
}

func GetComponentService(projectID int) ([]models.Component, error) {
	// Rule: the project must exist (a project with no components returns an empty list)
	active, err := repository.ProjectActiveRepo(projectID)
	if err != nil {
		return nil, err
	}
	if !active {
		return nil, ErrProjectNotFound
	}
	return repository.GetComponentRepo(projectID)
}

func GetByIdComponentService(id int) (models.Component, error) {
	component, err := repository.GetByIdComponentRepo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return component, ErrComponentNotFound
	}
	return component, err
}

func UpdateComponentService(id int, component models.Component) error {
	existing, err := repository.GetByIdComponentRepo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrComponentNotFound
	}
	if err != nil {
		return err
	}
	if err := cleanComponent(&component); err != nil {
		return err
	}

	if component.Status == nil {
		component.Status = existing.Status
	}

	exists, err := repository.ComponentNameExistsRepo(existing.ProjectId, component.ComponentName, id)
	if err != nil {
		return err
	}
	if exists {
		return ErrComponentExists
	}

	err = repository.UpdateComponentRepo(id, component)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrComponentNotFound
	}
	return err
}

func DeleteComponentService(id int) error {
	err := repository.DeleteComponentRepo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrComponentNotFound
	}
	return err
}

func cleanComponent(c *models.Component) error {
	c.ComponentName = strings.TrimSpace(c.ComponentName)
	c.Package = strings.TrimSpace(c.Package)
	c.PartNumber = strings.TrimSpace(c.PartNumber)
	c.Supplier = strings.TrimSpace(c.Supplier)

	if c.ComponentName == "" || c.Package == "" {
		return ErrInvalidInput
	}
	return nil
}
