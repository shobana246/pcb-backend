package models

type ProjectRequest struct {
	ProjectName string `json:"project_name" binding:"required,max=150"`
	Revision    string `json:"revision" binding:"required,max=100"`
	UpdatedBy   int    `json:"updated_by" binding:"required,gt=0"`
}

func (r ProjectRequest) ToModel() Project {
	return Project{
		ProjectName: r.ProjectName,
		Revision:    r.Revision,
		UpdatedBy:   r.UpdatedBy,
	}
}

type ComponentFields struct {
	ComponentName     string  `json:"component_name" binding:"required,max=100"`
	Package           string  `json:"package" binding:"required,max=250"`
	Category          string  `json:"category" binding:"required,oneof=Passive IC Connector Mechanical"`
	PlacementSide     string  `json:"placement_side" binding:"required,oneof=Top Bottom"`
	Rotation          float64 `json:"rotation" binding:"gte=0,lt=360"`
	Xposition         float64 `json:"x_position" binding:"gte=-999,lte=999"`
	Yposition         float64 `json:"y_position" binding:"gte=-999,lte=999"`
	Height            float64 `json:"height" binding:"gte=0,lte=99"`
	Supplier          string  `json:"supplier" binding:"max=250"`
	PartNumber        string  `json:"part_number" binding:"max=250"`
	Status            string  `json:"status" binding:"omitempty,oneof=verified warning critical unverified"`
	TolerancePosition float64 `json:"tolerance_position" binding:"gte=0,lte=99"`
	ToleranceRotation float64 `json:"tolerance_rotation" binding:"gte=0,lte=99"`
	ImageUrl          string  `json:"image_url" binding:"max=250"`
	Notes             *string `json:"notes" binding:"omitempty,max=255"`
}

func (f ComponentFields) ToModel() Component {
	category := ComponentCategory(f.Category)

	status := Unverified
	if f.Status != "" {
		status = ComponentStatus(f.Status)
	}

	return Component{
		ComponentName:     f.ComponentName,
		Package:           f.Package,
		Category:          &category,
		PlacementSide:     f.PlacementSide,
		Rotation:          f.Rotation,
		Xposition:         f.Xposition,
		Yposition:         f.Yposition,
		Height:            f.Height,
		Supplier:          f.Supplier,
		PartNumber:        f.PartNumber,
		Status:            &status,
		TolerancePosition: f.TolerancePosition,
		ToleranceRotation: f.ToleranceRotation,
		ImageUrl:          f.ImageUrl,
		Notes:             f.Notes,
	}
}

type CreateComponentRequest struct {
	ProjectID int `json:"project_id" binding:"required,gt=0"`
	ComponentFields
}

func (r CreateComponentRequest) ToModel() Component {
	c := r.ComponentFields.ToModel()
	c.ProjectId = r.ProjectID
	return c
}

type CreateIssueRequest struct {
	ComponentID int    `json:"component_id" binding:"required,gt=0"`
	Issue       string `json:"issue" binding:"required,max=255"`
	Severity    string `json:"severity" binding:"required,oneof=warning critical"`
}

func (r CreateIssueRequest) ToModel() Issue {
	return Issue{
		ComponentId: r.ComponentID,
		Issues:      r.Issue,
		Severity:    IssueSeverity(r.Severity),
		Status:      Open,
	}
}

type UpdateIssueRequest struct {
	Status string `json:"status" binding:"required,oneof=open 'in review' resolved"`
}
