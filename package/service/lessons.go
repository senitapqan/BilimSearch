package service

import "BilimSearch/models"

func (s *service) GetMyLessons(studentId int) ([]models.Lesson, error) {
	lessons, err := s.repos.GetMyLessons(studentId)
	if err != nil {
		return nil, err
	}
	return lessons, nil
}

func (s *service) GetMyLessonById(courseId, studentId int) (models.Lesson, error) {
	lesson, err := s.repos.GetMyLessonById(courseId, studentId)
	if err != nil {
		return lesson, err
	}
	return lesson, nil
}