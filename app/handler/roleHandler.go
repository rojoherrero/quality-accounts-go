package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"github.com/rojoherrero/quality-common"
	"net/http"

	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-accounts/app/service"
)

type (
	RoleHandler interface {
		Save(w http.ResponseWriter, r *http.Request)
		Update(w http.ResponseWriter, r *http.Request)
		Paginate(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
	}

	roleHandler struct {
		service service.RoleService
		logger  common.Logger
	}
)

func NewRoleHandler(service service.RoleService, logger common.Logger) RoleHandler {
	return &roleHandler{
		service: service,
		logger:  logger,
	}
}

func (h *roleHandler) Save(w http.ResponseWriter, r *http.Request) {
	role, _ := model.UnmarshalRoleDepartment(r.Body)
	if e := h.service.Save(role); e != nil {

	}
}

func (h *roleHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *roleHandler) Paginate(w http.ResponseWriter, r *http.Request) {

}

func (h *roleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if e := h.service.Delete(code); e != nil {

	}
}
