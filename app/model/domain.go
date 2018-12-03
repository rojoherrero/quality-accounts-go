package model

import (
	"encoding/json"
	"io"
	"time"
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

type Account struct {
	Username    string           `json:"username"`
	Password    string           `json:"password"`
	Roles       RolesDepartments `json:"roles"`
	Departments RolesDepartments `json:"departments"`
	FullName    string           `json:"fullName"`
	Creation    time.Time        `json:"creation"`
	Update      time.Time        `json:"update"`
}

type Accounts []Account

func UnmarshallAccount(data io.ReadCloser) (Account, error) {
	var a Account
	e := json.NewDecoder(data).Decode(&a)
	return a, e
}

func (a *Account) Marshall() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Accounts) Marshall() ([]byte, error) {
	return json.Marshal(a)
}
