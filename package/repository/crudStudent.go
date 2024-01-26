package repository

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
	"fmt"
)

func (r *repository) CreateStudent(student models.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	user_id, err := r.CreateUser(student, tx)
	
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("insert into %s (user_id) values($1) returning id", studentsTable)	
	row := tx.QueryRow(query, user_id)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r repository) DeleteStudent(lessonId int) error {
	return nil
}

func (r repository) GetStudent(lessonId int) (dtos.User, error) {
	var student dtos.User
	return student, nil
}

func (r repository) GetStudents() ([]dtos.User,  error) {
	return nil, nil
}

