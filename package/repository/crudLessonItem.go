package repository

import (
	"BilimSearch/models"
	"fmt"
)

func (r repository) CreateLessonItem(lessonItem models.LessonItem) (int, error) {
	var lessonItemId int
	request := fmt.Sprintf("insert into %s (lesson_id, date, topic) values ($1, $2, $3) returning id", lessonItemTable)
	row := r.db.QueryRowx(request, lessonItem.LessonId, lessonItem.Date, lessonItem.Topic)
	if err := row.Scan(&lessonItemId); err != nil {
		return -1, err
	}	
	return lessonItemId, nil
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