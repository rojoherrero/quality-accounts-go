package model

import (
	"encoding/json"
	"io"
)

type RoleDepartmentUpdate struct {
	NewCode        string             `json:"newCode"`
	NewDescription string             `json:"newDescription"`
	OldCode        string             `json:"oldCode"`
	Type           RoleDepartmentType `json:"type"`
}

func UnmarshallRoleDepartmentUpdate(data io.ReadCloser) (RoleDepartmentUpdate, error) {
	var r RoleDepartmentUpdate
	e := json.NewDecoder(data).Decode(&r)
	return r, e
}

func (du *RoleDepartmentUpdate) Marshall() ([]byte, error) {
	return json.Marshal(du)
}

