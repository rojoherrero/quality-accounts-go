package handler

import (
	"encoding/json"
	"github.com/rojoherrero/quality-accounts/app/model/entity"
	"github.com/rojoherrero/quality-accounts/app/service"
	"net/http"
)

type (
	RoleHandler interface {
		Save(res http.ResponseWriter, req *http.Request)
		Update(res http.ResponseWriter, req *http.Request)
		Paginate(res http.ResponseWriter, req *http.Request)
		Delete(res http.ResponseWriter, req *http.Request)
	}

	roleHandler struct {
		service service.RoleService
	}
)

func NewRoleHandler(service service.RoleService) RoleHandler {
	return &roleHandler{service: service}
}

func (h *roleHandler) Save(res http.ResponseWriter, req *http.Request) {
	var role entity.Role
	if e := json.NewDecoder(req.Body).Decode(&role); e != nil {

	}
	if e := h.service.Save(role); e != nil {

	}
}

func (h *roleHandler) Update(res http.ResponseWriter, req *http.Request) {

}

func (h *roleHandler) Paginate(res http.ResponseWriter, req *http.Request) {

}

func (h *roleHandler) Delete(res http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	if e := h.service.Delete(code); e != nil {

	}
}
