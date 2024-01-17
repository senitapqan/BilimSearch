package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) homeGrades(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	roleId, err := getRoleId(c, strudentCtx)

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

}

func (h *Handler) homeAttendance(c *gin.Context) {

}

func (h *Handler) attendanceItem(c *gin.Context) {
	
}