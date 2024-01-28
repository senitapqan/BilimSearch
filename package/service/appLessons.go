package service

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
)

func (s *service) GetMyLessons(studentId int) ([]models.Lesson, error) {
	lessons, err := s.repos.GetMyLessons(studentId)
	return lessons, err
}

func (s *service) GetMyLessonItemsById(courseId, studentId int) ([]dtos.LessonItemResponse, error) {
	lessonItems, err := s.repos.GetMyLessonItemsById(courseId, studentId)
	return lessonItems, err
}

