package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/rojoherrero/quality-common"

	"github.com/rojoherrero/quality-accounts/app/model"
	"github.com/rojoherrero/quality-accounts/app/service"
)

type (
	RoleHandler interface {
		Save(c *gin.Context)
		Update(c *gin.Context)
		Paginate(c *gin.Context)
		Delete(c *gin.Context)
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

func (h *roleHandler) Save(c *gin.Context) {
	var role model.RoleDepartment
	c.BindJSON(&role)
	if e := h.service.Save(role); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		h.logger.Error(e.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *roleHandler) Update(c *gin.Context) {
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

func (h *roleHandler) Paginate(c *gin.Context) {
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
	var roles model.RolesDepartments
	if roles, e = h.service.Paginate(start, end); e != nil {
		h.logger.Error(e.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *roleHandler) Delete(c *gin.Context) {
	code := c.Query("code")
	if e := h.service.Delete(code); e != nil {
		h.logger.Error(e.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
