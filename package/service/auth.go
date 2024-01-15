package service

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int      `json:"user_id"`
	Roles  []string `json:"roles"`
}

func (s *service) hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *service) GenerateToken(username, password string) (string, error) {
	password = s.hashPassword(password)
	user, err := s.repos.GetUser(username)

	if err != nil {
		return "",err
	}

	if user.Password != password {
		return "", http.ErrAbortHandler
	}

	roles, err := s.repos.GetRoles(user.Id)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		roles,
	})
	
	return token.SignedString([]byte(signingKey))
}
