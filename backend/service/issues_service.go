package service

import (
	"database/sql"
	"errors"
	"pcb-backend/backend/models"
	"pcb-backend/backend/repository"
	"strings"
)

var (
	ErrIssueNotFound     = errors.New("issue not found")
	ErrInvalidTransition = errors.New("status change not allowed")
)

var allowedNext = map[models.Status][]models.Status{
	models.Open:     {models.InReview},
	models.InReview: {models.Open, models.Resolved},
	models.Resolved: {models.InReview},
}

func canMove(from, to models.Status) bool {
	for _, s := range allowedNext[from] {
		if s == to {
			return true
		}
	}
	return false
}

func CreateIssuesService(issue models.Issue) error {
	issue.Issues = strings.TrimSpace(issue.Issues)
	if issue.Issues == "" {
		return ErrInvalidInput
	}

	active, err := repository.ComponentActiveRepo(issue.ComponentId)
	if err != nil {
		return err
	}
	if !active {
		return ErrComponentNotFound
	}

	return repository.CreateIssuesRepo(issue)
}

func GetAllIssuesService() ([]models.Issue, error) {
	return repository.GetAllIssuesRepo()
}

func GetIssuesByIdService(id int) (models.Issue, error) {
	issue, err := repository.GetIssuesByIdRepo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return issue, ErrIssueNotFound
	}
	return issue, err
}

func UpdateIssueByIdService(id int, status models.Status) error {
	issue, err := repository.GetIssuesByIdRepo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrIssueNotFound
	}
	if err != nil {
		return err
	}

	if !canMove(issue.Status, status) {
		return ErrInvalidTransition
	}

	err = repository.UpdateIssueByIdRepo(id, status)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrIssueNotFound
	}
	return err
}

func DeleteIssueService(id int) error {
	err := repository.DeleteIssueRepo(id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrIssueNotFound
	}
	return err
}
