package models

import "time"

type Project struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	ProjectName   string    `json:"project_name"`
	Revision      string    `json:"revision"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedBy     int       `json:"updated_by"`
	UpdatedByName string    `gorm:"->;-:migration" json:"updatedby"`
}

func (Project) TableName() string {
	return "projects"
}
