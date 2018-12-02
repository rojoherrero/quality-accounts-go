package model

import (
	"encoding/json"
	"io"
)

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

func UnmarshalRoleDepartment(data io.ReadCloser) (RoleDepartment, error) {
	var r RoleDepartment
	e := json.NewDecoder(data).Decode(&r)
	return r, e
}

func (r *RoleDepartment) Marshall() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RolesDepartments) Marshall() ([]byte, error) {
	return json.Marshal(r)
}
