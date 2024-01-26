package repository

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
)

func (r repository) CreateTeacher(teacher models.Teacher) (int, error) {
	return 0, nil
}

func (r repository) DeleteTeacher(teacherId int) error {
	return nil
}

func (r repository) GetTeacher(teacherId int) (dtos.User, error) {
	var teacher dtos.User
	return teacher, nil
}

func (r repository) GetTeachers() ([]dtos.User, error) {
	return nil, nil
}
