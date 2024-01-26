package service

import "BilimSearch/models"

func (s service) CreateLesson(lesson models.Lesson) (int, error) {
	return 0, nil
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