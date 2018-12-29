package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rojoherrero/quality-accounts/backend/model"
	"github.com/rojoherrero/quality-accounts/backend/service"
)

const (
	_10SecondsTimeOut = 10 * time.Second
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
	}
)

func NewRoleHandler(service service.RoleService) RoleHandler {
	return &roleHandler{service}
}

func (h *roleHandler) Save(c *gin.Context) {
	var role []model.Role
	ctx, _ := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	_ = c.BindJSON(&role)
	if e := h.service.Save(ctx, role); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *roleHandler) Update(c *gin.Context) {
	var update model.Role
	_ = c.BindJSON(&update)
	code := c.Query("code")
	ctx, _ := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	if e := h.service.Update(ctx, update, code); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *roleHandler) Paginate(c *gin.Context) {
	start, e := strconv.Atoi(c.Param("start"))
	if e != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	end, e := strconv.Atoi(c.Param("end"))
	if e != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx, _ := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	roles, e := h.service.Paginate(ctx, start, end)
	if e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *roleHandler) Delete(c *gin.Context) {
	code := c.Query("code")
	ctx, _ := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	if e := h.service.Delete(ctx, code); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
