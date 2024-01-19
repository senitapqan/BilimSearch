package repository

import (
	"BilimSearch/dtos"
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

func (r *repository) GetMyLessonItemsById(courseId, studentId int) ([]dtos.LessonItemResponse, error) {
	var lessonItems []dtos.LessonItemResponse
	query := fmt.Sprintf(`select li.id, li.date, t.id from %s li
							join %s t on li.id = t.lesson_item_id
							where li.lesson_id = $1
							order by li.id`, lessonItemTable, tasksTable)

	rows, err := r.db.Query(query, courseId)
	
	if err != nil {
		return nil, err
	}

	last := models.LessonItem{Id: -1}
	tasks := []models.Task{}
	var task models.Task
	var lessonItem models.LessonItem

	for rows.Next() {
		err := rows.Scan(&lessonItem.Id, &lessonItem.Date, &task.Id)
		if err != nil {
			return nil, err
		}

		if last.Id == -1 {
			last = lessonItem
		}

		if lessonItem.Id == last.Id {
			tasks = append(tasks, task)
		} else {
			lessonItems = append(lessonItems, dtos.LessonItemResponse{LessonItem: last, Tasks: tasks})
			tasks = []models.Task{}
			last = lessonItem
		}
	}
	if last.Id != -1 {
		lessonItems = append(lessonItems, dtos.LessonItemResponse{LessonItem: lessonItem, Tasks: tasks})
		tasks = []models.Task{}
	}
	return lessonItems, err
}