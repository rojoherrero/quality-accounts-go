package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"net/http"

	"github.com/rojoherrero/quality-accounts/app/service"
)

type (
	DepartmentHandler interface {
		Save(w http.ResponseWriter, r *http.Request)
		Update(w http.ResponseWriter, r *http.Request)
		Paginate(w http.ResponseWriter, r *http.Request)
		Delete(w http.ResponseWriter, r *http.Request)
	}

	departmentHandler struct {
		service service.DepartmentService
	}
)

func NewDepartmentHandler(service service.DepartmentService) DepartmentHandler {
	return &departmentHandler{service: service}
}

func (h *departmentHandler) Save(w http.ResponseWriter, r *http.Request) {}
func (h *departmentHandler) Update(w http.ResponseWriter, r *http.Request) {}
func (h *departmentHandler) Paginate(w http.ResponseWriter, r *http.Request) {}
func (h *departmentHandler) Delete(w http.ResponseWriter, r *http.Request) {}
