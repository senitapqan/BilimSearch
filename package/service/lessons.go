package service

import "BilimSearch/models"

func (s *service) GetMyLessons(studentId int) ([]models.Lesson, error) {
	lessons, err := s.repos.GetMyLessons(studentId)
	if err != nil {
		return nil, err
	}
	return lessons, nil
}