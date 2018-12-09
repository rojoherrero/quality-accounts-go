package model

import (
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
