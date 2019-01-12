package model

type UserCreationDto struct {
	ID              int64           `json:"id,omitempty"`
	FullName        string          `json:"fullName"`
	UserName        string          `json:"userName"`
	Password        string          `json:"password"`
	DepartmentRoles DepartmentRoles `json:"departmentRoles"`
}

type DepartmentRoles struct {
	DepartmentCode string   `json:"departmentCode"`
	RolesCodes     []string `json:"rolesCodes"`
}
