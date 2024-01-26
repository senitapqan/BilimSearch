package service

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
)

func (s *service) GetMyCourses(studentId int) ([]models.Course, error) {
	courses, err := s.repos.GetMyCourses(studentId)

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (s *service) GetMyCoursesGrades(studentId int) ([]dtos.CourseGrades, error) {
	grades, err := s.repos.GetMyCoursesGrades(studentId)

	if err != nil {
		return nil, err
	}
	return grades, err
}