package models

import "time"

type ComponentCategory string

const (
	Passive    ComponentCategory = "Passive"
	IC         ComponentCategory = "IC"
	Connector  ComponentCategory = "Connector"
	Mechanical ComponentCategory = "Mechanical"
)

type ComponentStatus string

const (
	Verified   ComponentStatus = "verified"
	Warning    ComponentStatus = "warning"
	Critical   ComponentStatus = "critical"
	Unverified ComponentStatus = "unverified"
)

type Component struct {
	Id                int                `gorm:"primaryKey" json:"id"`
	ProjectId         int                `json:"project_id"`
	ComponentName     string             `json:"component_name"`
	Package           string             `json:"package"`
	Category          *ComponentCategory `json:"category"`
	PlacementSide     string             `json:"placement_side"`
	Rotation          float64            `json:"rotation"`
	Xposition         float64            `gorm:"column:x_position" json:"x_position"`
	Yposition         float64            `gorm:"column:y_position" json:"y_position"`
	Height            float64            `json:"height"`
	Supplier          string             `json:"supplier"`
	PartNumber        string             `json:"part_number"`
	Status            *ComponentStatus   `json:"status"`
	TolerancePosition float64            `json:"tolerance_position"`
	ToleranceRotation float64            `json:"tolerance_rotation"`
	ImageUrl          string             `json:"image_url"`
	Notes             *string            `json:"notes"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

func (Component) TableName() string {
	return "component"
}
