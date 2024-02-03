package service

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
	
)

func (s service) CreateTeacher(teacher models.User) (int, error) {
	teacher.Password = s.hashPassword(teacher.Password) 
	
	teacherId, err := s.repos.CreateTeacher(teacher)
	return teacherId, err
}

func (s service) DeleteTeacher(teacherId int) error {
	return s.repos.DeleteTeacher(teacherId)
}

func (s service) GetTeacher(teacherId int) (dtos.User, error) {
	var teacher dtos.User
	return teacher, nil
}

func (s service) GetTeachers() ([]dtos.User, error) {
	return nil, nil
}