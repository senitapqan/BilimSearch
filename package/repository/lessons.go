package repository

import (
	"BilimSearch/models"
	"fmt"

)

func (r *repository) GetMyLessons(studentId int) ([]models.Lesson, error) {
	var lessons []models.Lesson
	query := fmt.Sprintf(`select l.id, l.course_id, l.teacher_id, l.period from %s l
							join %s ls on l.id = ls.lesson_id 
							where ls.student_id = $1`, lessonsTable, lessonStudentsTable)
	rows, err := r.db.Query(query, studentId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var lesson models.Lesson
		err := rows.Scan(&lesson.Id, &lesson.CourseId, &lesson.TeacherId, &lesson.Period)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}