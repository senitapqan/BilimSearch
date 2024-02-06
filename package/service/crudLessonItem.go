package service

import "BilimSearch/models"

func (s service) CreateLessonItem(lessonItem models.LessonItem) (int, error) {
	lessonItemId, err := s.repos.CreateLessonItem(lessonItem)
	return lessonItemId, err
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