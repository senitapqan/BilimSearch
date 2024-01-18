package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) myCourse(c *gin.Context) {
	_, err := getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	roleId, err := getRoleId(c, strudentCtx)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}	

	courses, err := h.services.GetMyCourses(roleId)

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) myCourseItem(c *gin.Context) {
	
}

func (h *Handler) myCourseGrades(c *gin.Context) {
	
}

func (h *Handler) myCourseTasks(c *gin.Context) {
	
}

func (h *Handler) myCourseTaskItem(c *gin.Context) {
	
}

func (h *Handler) myCourseTaskItemSend(c *gin.Context) {
	
}