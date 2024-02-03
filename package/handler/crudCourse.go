package handler

import (
	"BilimSearch/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addCourse(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}

	var course models.Course
	if err := c.BindJSON(&course); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}	

	courseId, err := h.services.CreateCourse(course)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"course_id": courseId,
	})
}

func (h *Handler) deleteCourse(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}	

	courseId, err := h.getCourseId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DeleteCourse(courseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "course with id - " + strconv.Itoa(courseId) + " was deleted succesfully",
	})
}

func (h *Handler) getCourses(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}
}

func (h *Handler) getCourse(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}
}