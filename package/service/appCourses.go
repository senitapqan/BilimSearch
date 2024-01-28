package service

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
)

func (s *service) GetMyCourses(studentId int) ([]models.Course, error) {
	courses, err := s.repos.GetMyCourses(studentId)
	return courses, err
}

func (s *service) GetMyCoursesGrades(studentId int) ([]dtos.CourseGrades, error) {
	grades, err := s.repos.GetMyCoursesGrades(studentId)
	return grades, err
}