package service

import "BilimSearch/models"

func (s service) CreateCourse(course models.Course) (int, error){
	courseId, err := s.repos.CreateCourse(course)
	return courseId, err
}

func (s service) DeleteCourse(courseId int) error {
	return s.repos.DeleteCourse(courseId) 
}

func (s service) GetCourse(courseId int) (models.Course, error) {
	var course models.Course
	return course, nil
}

func (s service) GetCourses() ([]models.Course, error) {
	return nil, nil
}