package handler

import (
	"BilimSearch/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInRequest struct {
	username string
	password string
}

func (h *Handler) signIn(c *gin.Context) {
	var request signInRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	c.JSON(http.StatusOK, request)
}

func (h *Handler) signUp(c *gin.Context) {
	var request models.Student
	if err := c.BindJSON(&request); err != nil {
		return
	}

	id, err := h.services.CreateStudent(request)

	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]int{
		"new user was succesfully added with id": id,
	})
}