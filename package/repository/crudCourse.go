package repository

import (
	"BilimSearch/models"
	"fmt"
	"log"
)

func (r repository) CreateCourse(course models.Course) (int, error) {
	var courseId int
	query := fmt.Sprintf("insert into %s (name) values($1) returning id", coursesTable)
	row := r.db.QueryRowx(query, course.Name)
	if err := row.Scan(&courseId); err != nil {
		log.Printf("Error was after query: " + err.Error())
		return 0, err
	}
	return courseId, nil
}
func (r repository) DeleteCourse(courseId int) error {
	return nil
}
func (r repository) GetCourse(courseId int) (models.Course, error) {
	var course models.Course
	return course, nil
}
func (r repository) GetCourses() ([]models.Course, error) {
	return nil, nil
}
