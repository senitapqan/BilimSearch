package repository

import (
	"BilimSearch/models"
	"fmt"
	"log"

)

func (r repository) CreateLesson(lesson models.Lesson) (int, error) {
	var lessonId int
	query := fmt.Sprintf(`insert into %s (course_id, teacher_id, period) values ($1, $2, $3) returning id`, lessonsTable)
	row := r.db.QueryRowx(query, lesson.CourseId, lesson.TeacherId, lesson.Period)
	if err := row.Scan(&lessonId); err != nil {
		log.Printf("error was after query" + err.Error())
		return 0, err
	}
	return lessonId, nil
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

