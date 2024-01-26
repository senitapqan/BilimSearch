package service

import "BilimSearch/models"

func (s service) CreateLessonItem(lessonItem models.LessonItem) (int, error) {
	return 0, nil
}

func (s service) DeleteLessonItem(lessonItemId int) error {
	return nil
}

func (s service) GetLessonItem(lessonItemId int) (models.LessonItem, error) {
	var lessonItem models.LessonItem
	return lessonItem, nil
}

func (s service) GetLessonItems() ([]models.LessonItem, error) {
	return nil, nil
}