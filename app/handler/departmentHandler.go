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
	if e := h.service.Save(role); e != nil {
		h.logger.Error(common.JSON(w, http.StatusInternalServerError, nil).Error())
	}
	common.JSON(w, http.StatusOK, nil)
}

func (h *departmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	update, _ := model.UnmarshallRoleDepartmentUpdate(r.Body)
	if e := h.service.Update(update); e != nil{
		h.logger.Error(common.JSON(w, http.StatusInternalServerError, nil).Error())
	}
	common.JSON(w, http.StatusOK, nil)
}

func (h *departmentHandler) Paginate(w http.ResponseWriter, r *http.Request) {
	var e error
	var start int
	var end int
	vars := mux.Vars(r)
	if start, e = strconv.Atoi(vars["start"]); e != nil {
		h.logger.Error(common.JSON(w, http.StatusBadRequest, nil).Error())
	}
	if end, e = strconv.Atoi(vars["end"]); e != nil {
		h.logger.Error(common.JSON(w, http.StatusBadRequest, nil).Error())
	}
	var deps model.RolesDepartments
	if deps, e = h.service.Paginate(start, end); e != nil {
		h.logger.Error(common.JSON(w, http.StatusInternalServerError, nil).Error())
	}
	var bytes []byte
	if bytes, e = deps.Marshall(); e != nil {
		h.logger.Error(common.JSON(w, http.StatusInternalServerError, nil).Error())
	}
	common.JSON(w, http.StatusOK, bytes)
}

func (h *departmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if e := h.service.Delete(code); e != nil{
		h.logger.Error(common.JSON(w, http.StatusInternalServerError, nil).Error())
		return
	}
	common.JSON(w, http.StatusOK, nil)
}
