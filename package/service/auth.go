package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	Roles []string `json:"roles"`
}

func (s *service) hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *service) GenerateToken(username, password string) (int, string, error) {
	return 0, "", nil
}