package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "hello world",
	})
}

func (h *Handler) signUp(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "World Hello",
	})
}