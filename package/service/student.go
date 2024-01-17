package service

import (
	"BilimSearch/models"
)

func (s *service) CreateStudent(student models.User) (int, error) {
	student.Password = s.hashPassword(student.Password)
	student_id, err := s.repos.CreateStudent(student)
	if err != nil {
		return -1, err
	}
	return student_id, nil
}