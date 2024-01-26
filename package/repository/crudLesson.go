package repository

import "BilimSearch/models"

func (r repository) CreateLesson(lesson models.Lesson) (int, error) {
	return 0, nil
}

func (r repository) DeleteLesson(lessonId int) error {
	return nil
}

func (r repository) GetLesson(lessonId int) (models.Lesson, error) {
	var lesson models.Lesson
	return lesson, nil
}

func (r repository) GetLessons() ([]models.Lesson, error) {
	return nil, nil
}