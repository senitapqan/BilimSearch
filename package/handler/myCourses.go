package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) myCourse(c *gin.Context) {
	_, roleId, err := h.getIds(strudentCtx, c)

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
	_, roleId, err := h.getIds(strudentCtx, c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}	
	
	courseIdString := c.Param("id")
	if len(courseIdString) == 0 {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	courseId, err := strconv.Atoi(courseIdString)
	
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lesson, err := h.services.GetMyLessonById(courseId, roleId); 

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, lesson);
}

func (h *Handler) myCourseGrades(c *gin.Context) {
	
}

func (h *Handler) myCourseTasks(c *gin.Context) {
	
}

func (h *Handler) myCourseTaskItem(c *gin.Context) {
	
}

func (h *Handler) myCourseTaskItemSend(c *gin.Context) {
	
}