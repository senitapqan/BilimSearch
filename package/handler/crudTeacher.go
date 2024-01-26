package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addTeacher(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
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