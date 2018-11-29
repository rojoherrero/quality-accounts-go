package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-accounts/app/service"
	"github.com/rojoherrero/quality-common"
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
		logger  common.Logger
	}
)

func NewDepartmentHandler(service service.DepartmentService, logger common.Logger) DepartmentHandler {
	return &departmentHandler{
		service: service,
		logger:  logger,
	}
}

func (h *departmentHandler) Save(w http.ResponseWriter, r *http.Request) {
	role, _ := model.UnmarshalRoleDepartment(r.Body)
	h.service.Save(role)
	common.JSON(w, http.StatusOK, nil)
}

func (h *departmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	update, _ := model.UnmarshallRoleDepartmentUpdate(r.Body)
	h.service.Update(update)
	common.JSON(w, http.StatusOK, nil)
}

func (h *departmentHandler) Paginate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	start, e := strconv.Atoi(vars["start"])
	if e != nil {
		common.JSON(w, http.StatusBadRequest, nil)
	}
	end, _ := strconv.Atoi(vars["end"])
	if e != nil {
		common.JSON(w, http.StatusBadRequest, nil)
	}
	deps, _ := h.service.Paginate(start, end)
	common.JSON(w, http.StatusOK, deps)
}

func (h *departmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	h.service.Delete(code)
	common.JSON(w, http.StatusOK, nil)
}
