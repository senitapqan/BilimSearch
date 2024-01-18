package repository

import (
	"BilimSearch/models"
	"fmt"
)

func (r *repository) GetMyCourses(studentId int) ([]models.Course, error) {
	var courses []models.Course
	
	query := fmt.Sprintf(`select c.id, c.name from %s c 
							join %s l on c.id = l.course_id 
							join %s ls on l.id = ls.lesson_id 
							where ls.student_id = $1`, coursesTable, lessonsTable, lessonStudentsTable)

	rows, err := r.db.Query(query, studentId)

	if err != nil {
		return nil, err
	}							
	
	for rows.Next() {
		var course models.Course
		err := rows.Scan(&course.Id, &course.Name)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}