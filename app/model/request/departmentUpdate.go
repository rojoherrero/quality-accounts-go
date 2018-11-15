package request

type DepartmentUpdate struct {
	NewCode        string `json:"newCode"`
	NewDescription string `json:"newDescription"`
	OldCode        string `json:"oldCode"`
}
