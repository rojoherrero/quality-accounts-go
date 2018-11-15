package request

type RoleUpdate struct {
	NewCode        string `json:"newCode"`
	NewDescription string `json:"newDescription"`
	OldCode        string `json:"oldCode"`
}
