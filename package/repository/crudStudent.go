package repository

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
	"fmt"
)

func (r *repository) CreateStudent(student models.User) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var studentId int
	userId, err := r.CreateUser(student, tx)
	
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("insert into %s (user_id) values($1) returning id", studentsTable)	
	row := tx.QueryRowx(query, userId)

	if err := row.Scan(&studentId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return studentId, tx.Commit()
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

