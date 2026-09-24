package models

import "time"

type IssueSeverity string

const (
	SeverityWarning  IssueSeverity = "warning"
	SeverityCritical IssueSeverity = "critical"
)

type Status string

const (
	Open     Status = "open"
	InReview Status = "in review"
	Resolved Status = "resolved"
)

type Issue struct {
	ID          int           `gorm:"primaryKey" json:"id"`
	ComponentId int           `json:"component_id"`
	Reference   string        `gorm:"->;-:migration" json:"reference"`
	Issues      string        `gorm:"column:issue" json:"issue"`
	Severity    IssueSeverity `json:"severity"`
	Status      Status        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
}

func (Issue) TableName() string {
	return "placement_issues"
}
