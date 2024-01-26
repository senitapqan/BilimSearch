package service

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
)

func (s service) CreateTeacher(teacher models.Teacher) (int, error) {
	return 0, nil
}

func (s service) DeleteTeacher(teacherId int) error {
	return nil
}

func (s service) GetTeacher(teacherId int) (dtos.User, error) {
	var teacher dtos.User
	return teacher, nil
}

func (s service) GetTeachers() ([]dtos.User, error) {
	return nil, nil
}