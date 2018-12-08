package model

type RoleDepartmentType string

const (
	Role       RoleDepartmentType = "Role"
	Department RoleDepartmentType = "Department"
)

type RoleDepartment struct {
	Code        string             `json:"code"`
	Description string             `json:"description"`
	Type        RoleDepartmentType `json:"type"`
}

type RolesDepartments []RoleDepartment

