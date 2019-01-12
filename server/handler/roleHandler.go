package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rojoherrero/quality-accounts/server/model"
	"github.com/rojoherrero/quality-accounts/server/service"
)

const (
	_10SecondsTimeOut = 100 * time.Hour
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
		logger  zerolog.Logger
	}
)

func NewRoleHandler(service service.RoleService, logger zerolog.Logger) RoleHandler {
	return &roleHandler{service, logger}
}

func (h *roleHandler) Save(c *gin.Context) {
	var role []model.Role
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	if e := h.service.Update(ctx, update); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *roleHandler) Paginate(c *gin.Context) {
	start, e := strconv.Atoi(c.Query("start"))
	if e != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	end, e := strconv.Atoi(c.Query("end"))
	if e != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	roles, e := h.service.Paginate(ctx, start, end)
	if e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *roleHandler) Delete(c *gin.Context) {
	code := c.Query("code")
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	if e := h.service.Delete(ctx, code); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
