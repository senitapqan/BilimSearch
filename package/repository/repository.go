package repository

import (
	"BilimSearch/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUser(username string) (models.User, error)
	GetRoles(id int) ([]string, error)

	CreateStudent(student models.Student) (int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}