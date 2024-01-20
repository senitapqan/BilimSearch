package service

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
	"BilimSearch/package/repository"
)

type Service interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(tokem string) (int, []models.RolesHeaders, error)
	
	GetMyCourses(studentId int) ([]models.Course, error)
	GetMyCoursesGrades(studentId int) ([]dtos.CourseGrades, error)

	GetMyLessons(studentId int) ([]models.Lesson, error)
	GetMyLessonItemsById(courseId, studentId int) ([]dtos.LessonItemResponse, error)

	CreateStudent(student models.User) (int, error)
}

type service struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) Service {
	return &service{
		repos: repos,
	}
}