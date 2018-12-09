package model

type RoleDepartmentUpdate struct {
	NewCode        string             `json:"newCode"`
	NewDescription string             `json:"newDescription"`
	OldCode        string             `json:"oldCode"`
	Type           RoleDepartmentType `json:"type"`
}



