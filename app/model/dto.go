package model

type UserCreationDto struct {
	ID              int64               `json:"id"`
	FullName        string              `json:"fullName"`
	UserName        string              `json:"userName"`
	Password        string              `json:"password"`
	DepartmentRoles map[string][]string `json:"departmentRoles"`
}
