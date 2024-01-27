package handler

import (
	"BilimSearch/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addTeacher(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var teacher models.User
	if err := c.BindJSON(&teacher); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} 

	teacherId, err := h.services.CreateTeacher(teacher)
	if  err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"teacher_id": teacherId,
	})
}

func (h *Handler) deleteTeacher(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getTeachers(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getTeacher(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}