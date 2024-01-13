package repository

import (
	"BilimSearch/models"
	"fmt"
)

func (r *repository) CreateStudent(student models.Student) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	user_id, err := r.CreateUser(student.Username, student.Password, tx)
	
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("insert into %s (username, password, email, name, surname, user_id) values($1, $2, $3, $4, $5, $6) returning id", studentsTable)	
	row := tx.QueryRow(query, student.Username, student.Password, student.Email, student.Name, student.Surname, user_id)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}