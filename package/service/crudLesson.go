package service

import "BilimSearch/models"

func (s service) CreateLesson(lesson models.Lesson) (int, error) {
	lessonId, err := s.repos.CreateLesson(lesson)
	return lessonId, err;
}

func (s service) DeleteLesson(lessonId int) error {
	return nil
}

func (s service) GetLesson(lessonId int) (models.Lesson, error) {
	var lesson models.Lesson
	return lesson, nil
}

func (s service) GetLessons() ([]models.Lesson, error) {
	return nil, nil
}