package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) getCourseId(c *gin.Context) (int, error){
	courseIdString := c.Param("id")
	if len(courseIdString) == 0 {
		return -1, errors.New("Can't find courseId in URL")
	}

	courseId, err := strconv.Atoi(courseIdString)
	
	if err != nil {
		return -1, err
	}		
	return courseId, nil
}

func (h *Handler) myCourse(c *gin.Context) {
	_, roleId, err := h.getIds(strudentCtx, c)

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
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
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}
	
	courseId, err := h.	getCourseId(c)
	
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lessonItems, err := h.services.GetMyLessonItemsById(courseId, roleId); 

	if err != nil {
		newErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	c.JSON(http.StatusOK, lessonItems);
}

func (h *Handler) myCourseGrades(c *gin.Context) {
	
}

func (h *Handler) myCourseTask(c *gin.Context) {
	
}

func (h *Handler) myCourseTaskSend(c *gin.Context) {
	
}