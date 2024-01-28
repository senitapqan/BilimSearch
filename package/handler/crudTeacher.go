package handler

import (
	"BilimSearch/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getTeacherId(c *gin.Context) (int, error){
	teacherIdString := c.Param("id")
	if len(teacherIdString) == 0 {
		return -1, errors.New("Can't find courseId in URL")
	}

	teacherId, err := strconv.Atoi(teacherIdString)
	
	if err != nil {
		return -1, err
	}		
	return teacherId, nil
}

func (h *Handler) addTeacher(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}

	var teacher models.User
	if err := c.BindJSON(&teacher); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} 

	teacherId, err := h.services.CreateTeacher(teacher)
	if  err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"teacher_id": teacherId,
	})
}

func (h *Handler) deleteTeacher(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}

	teacherId, err := h.getTeacherId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DeleteTeacher(teacherId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "teacher with id - " + strconv.Itoa(teacherId) + " was deleted succesfully",
	})
}

func (h *Handler) getTeachers(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}
}

func (h *Handler) getTeacher(c *gin.Context) {
	_, _, err := h.getIds(adminCtx, c)	

	if err != nil {
		newErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		return
	}
}