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
	DeleteStudent(studentId int) error
	GetStudent(studentId int) (dtos.User, error)
	GetStudents() ([]dtos.User, error)

	CreateTeacher(teacher models.Teacher) (int, error)
	DeleteTeacher(teacherId int) error
	GetTeacher(teacherId int) (dtos.User, error)
	GetTeachers() ([]dtos.User, error)

	CreateCourse(course models.Course) (int, error)
	DeleteCourse(courseId int) error
	GetCourse(courseId int) (models.Course, error)
	GetCourses() ([]models.Course, error)

	CreateLesson(lesson models.Lesson) (int, error)
	DeleteLesson(lessonId int) error
	GetLesson(lessonId int) (models.Lesson, error)
	GetLessons() ([]models.Lesson, error)

	CreateLessonItem(lessonItem models.LessonItem) (int, error)
	DeleteLessonItem(lessonItemId int) error
	GetLessonItem(lessonItemId int) (models.LessonItem, error)
	GetLessonItems() ([]models.LessonItem, error)
}

type service struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) Service {
	return &service{
		repos: repos,
	}
}