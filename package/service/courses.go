package service

import "BilimSearch/models"

func (s *service) GetMyCourses(studentId int) ([]models.Course, error) {
	courses, err := s.repos.GetMyCourses(studentId)

	if err != nil {
		return nil, err
	}

	return courses, nil
}