package repository

import "BilimSearch/models"

func (r repository) CreateCourse(course models.Course) (int, error) {
	return 0, nil
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