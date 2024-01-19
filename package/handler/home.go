package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) homeGrades(c *gin.Context) {
	userId, roleId, err := h.getIds(strudentCtx, c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, map[string]string {
		"userId": strconv.Itoa(userId), 
		"student": strconv.Itoa(roleId),
	})
}

func (h *Handler) homeSchedule(c *gin.Context) {
	_, roleId, err := h.getIds(strudentCtx, c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lessons, err := h.services.GetMyLessons(roleId)

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, lessons)
}

func (h *Handler) homeAttendance(c *gin.Context) {
	
}

func (h *Handler) attendanceItem(c *gin.Context) {
	
}