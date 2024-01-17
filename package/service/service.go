package service

import (
	"BilimSearch/models"
	"BilimSearch/package/repository"
)

type Service interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(tokem string) (int, []models.RolesHeaders, error)
	
	CreateStudent(student models.User) (int, error)
}

type service struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) Service {
	return &service{
		repos: repos,
	}
}