package service

import (
	"BilimSearch/models"
	"BilimSearch/package/repository"
)

type Service interface {
	CreateStudent(student models.Student) (int, error)
}

type service struct {
	repos *repository.Repository
}

func NewService(repos *repository.Repository) Service {
	return &service{
		repos: repos,
	}
}