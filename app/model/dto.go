package model

type UserCreationDto struct {
	FullName         string              `json:"fullName"`
	UserName         string              `json:"userName"`
	Password         string              `json:"password"`
	RolesDepartments map[string][]string `json:"rolesDepartments"`
}
