package handler

import (
	"BilimSearch/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addLessonItem(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}

	var lessonItem models.LessonItem
	if err := c.BindJSON(&lessonItem); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lessonItemId, err := h.services.CreateLessonItem(lessonItem)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int {
		"message": lessonItemId,
	})
}

func (h *Handler) deleteLessonItem(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}
}

func (h *Handler) getLessonItems(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}	
}

func (h *Handler) getLessonItem(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}
}