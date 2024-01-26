package repository

import "BilimSearch/models"

func (r repository) CreateLessonItem(lessonItem models.LessonItem) (int, error) {
	return 0, nil
}

func (r repository) DeleteLessonItem(lessonItemId int) error {
	return nil
}

func (r repository) GetLessonItem(lessonItemId int) (models.LessonItem, error) {
	var lessonItem models.LessonItem
	return lessonItem, nil
}

func (r repository) GetLessonItems() ([]models.LessonItem, error) {
	return nil, nil
}