package repository

import (
	"BilimSearch/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUser(username string) (models.User, error)
	GetRoles(id int) ([]string, error)

	GetId(role string, id int) (int, error)

	GetMyCourses(id int) ([]models.Course, error)

	GetMyLessons(id int) ([]models.Lesson, error)

	CreateStudent(student models.User) (int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}