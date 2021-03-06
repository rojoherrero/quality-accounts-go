package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rojoherrero/quality-accounts/server/model"
	"github.com/rojoherrero/quality-accounts/server/service"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type (
	UserHandler interface {
		Save(c *gin.Context)
		Update(c *gin.Context)
		Paginate(c *gin.Context)
		Delete(c *gin.Context)
	}

	userHandler struct {
		service service.UserService
		logger  zerolog.Logger
	}
)

func NewUserHandler(service service.UserService, logger zerolog.Logger) UserHandler {
	return &userHandler{service: service, logger: logger}
}

func (h *userHandler) Save(c *gin.Context) {
	var user model.UserCreationDto
	_ = c.BindJSON(&user)
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	if e := h.service.Save(ctx, user); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *userHandler) Update(c *gin.Context) {
	var user model.UserCreationDto
	_ = c.BindJSON(&user)
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	if e := h.service.Update(ctx, user); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *userHandler) Paginate(c *gin.Context) {
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
	roles, e := h.service.Paginate(ctx, int64(start), int64(end))
	if e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *userHandler) Delete(c *gin.Context) {
	userId, e := strconv.Atoi(c.Query("userId"))
	if e != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	if e := h.service.Delete(ctx, int64(userId)); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
