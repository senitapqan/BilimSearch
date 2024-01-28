package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addLessonItem(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}
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