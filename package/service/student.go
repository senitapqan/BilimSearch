package service

import (
	"BilimSearch/models"
)

func (s *service) CreateStudent(student models.Student) (int, error) {
	student.Password = s.hashPassword(student.Password)
	student_id, err := s.repos.CreateStudent(student)
	if err != nil {
		return -1, err
	}
	return student_id, nil
}