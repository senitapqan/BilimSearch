package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	strudentCtx         = "STUDENT"
	teacherCtx          = "TEACHER"
	adminCtx 			= "ADMIN"
)

func (h *Handler) userIdentify() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			newErrorResponse(c, http.StatusUnauthorized, "invalid header")
			return
		}

		if len(headerParts[1]) == 0 {
			newErrorResponse(c, http.StatusUnauthorized, "token is empty")
			return
		}

		userId, roles, err := h.services.ParseToken(headerParts[1])

		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, "Parsing works wrong: " + err.Error())
			return
		}

		c.Request.Header.Add(userCtx, strconv.Itoa(userId))
		for _, role := range roles {
			c.Request.Header.Add(role.Role, strconv.Itoa(role.Id))
		}

		c.Next()
	}
}

func getUserId(c *gin.Context) (int, error) {
	id := c.GetHeader(userCtx)
	if id == "" {
		return 0, errors.New("user id not found")
	}

	var userId int
	userId, err := strconv.Atoi(id)

	if err != nil {
		return 0, errors.New("cant converse user id")
	}

	return userId, nil
}

func getRoleId(c *gin.Context, role string) (int, error) {
	id := c.GetHeader(role)
	if id == "" {
		return 0, errors.New("role id not found")
	}

	var roleId int
	roleId, err := strconv.Atoi(id)

	if err != nil {
		return 0, errors.New("cant converse role id")
	}

	return roleId, nil	
}