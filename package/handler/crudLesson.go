package handler

import (
	"BilimSearch/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addLesson(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var lesson models.Lesson
	if err := c.BindJSON(&lesson); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lessonId, err := h.services.CreateLesson(lesson)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "created a new lesson with id " + strconv.Itoa(lessonId),
	})
}

func (h *Handler) deleteLesson(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getLessons(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getLesson(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}