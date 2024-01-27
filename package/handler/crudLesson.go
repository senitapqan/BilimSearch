package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addLesson(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
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