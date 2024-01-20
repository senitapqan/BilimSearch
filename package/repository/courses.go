package repository

import (
	"BilimSearch/dtos"
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

func (r *repository) GetMyCoursesGrades(studentId int) ([]dtos.CourseGrades, error) {
	var grades []dtos.CourseGrades
	query := fmt.Sprintf(`select c.name, g.prefinal_grade, g.final_grade from %s g
							join %s c on c.id = g.course_id
							join %s s on s.id = g.student_id
							where s.id = $1`, coursesStudentsGradesTable, coursesTable, studentsTable)
	rows, err := r.db.Queryx(query, studentId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var grade dtos.CourseGrades
		err := rows.Scan(&grade.CourseName, &grade.PrefinalGrade, &grade.FinalGrade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}	
	return grades, nil
}