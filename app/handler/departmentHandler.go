package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-accounts/app/service"
	"github.com/rojoherrero/quality-common"
)

type (
	DepartmentHandler interface {
		Save(c *gin.Context)
		Update(c *gin.Context)
		Paginate(c *gin.Context)
		Delete(c *gin.Context)
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

func (h *departmentHandler) Save(c *gin.Context) {
	var dep model.RoleDepartment
	_ = c.BindJSON(&dep)
	if e := h.service.Save(dep); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		h.logger.Error(e.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *departmentHandler) Update(c *gin.Context) {
	var update model.RoleDepartment
	c.BindJSON(&update)
	code := c.Query("code")
	if e := h.service.Update(update, code); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		h.logger.Error(e.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *departmentHandler) Paginate(c *gin.Context) {
	var e error
	var start int
	var end int
	if start, e = strconv.Atoi(c.Param("start")); e != nil {
		h.logger.Error(e.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if end, e = strconv.Atoi(c.Param("end")); e != nil {
		h.logger.Error(e.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	var deps model.RolesDepartments
	if deps, e = h.service.Paginate(start, end); e != nil {
		h.logger.Error(e.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, deps)
}

func (h *departmentHandler) Delete(c *gin.Context) {
	code := c.Query("code")
	if e := h.service.Delete(code); e != nil {
		h.logger.Error(e.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
