package handler

//go:generate mockgen -source=$GOFILE -destination=../mock/mock_$GOFILE -package=mock

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rojoherrero/quality-accounts/backend/model"
	"github.com/rojoherrero/quality-accounts/backend/service"
	"net/http"
	"strconv"
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
	}
)

func NewDepartmentHandler(service service.DepartmentService) DepartmentHandler {
	return &departmentHandler{service: service}
}

func (h *departmentHandler) Save(c *gin.Context) {
	var departments []model.Department
	_ = c.BindJSON(&departments)
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	if e := h.service.Save(ctx, departments); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *departmentHandler) Update(c *gin.Context) {
	var department model.Department
	c.BindJSON(&department)
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	if e := h.service.Update(ctx, department); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *departmentHandler) Paginate(c *gin.Context) {
	var e error
	var start int
	var end int
	if start, e = strconv.Atoi(c.Query("start")); e != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if end, e = strconv.Atoi(c.Query("end")); e != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	var deps []model.Department
	if deps, e = h.service.Paginate(ctx, start, end); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, deps)
}

func (h *departmentHandler) Delete(c *gin.Context) {
	id := c.Query("code")
	ctx, cancel := context.WithTimeout(c.Request.Context(), _10SecondsTimeOut)
	defer cancel()
	if e := h.service.Delete(ctx, id); e != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
