package repository

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
	"fmt"
)

func (r repository) CreateTeacher(teacher models.User) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	userId, err := r.CreateUser(teacher, tx)

	if err != nil {
		return 0, err
	}

	var teacherId int
	query := fmt.Sprintf(`insert into %s (user_id) values ($1) returning id`, teachersTable)
	row := tx.QueryRowx(query, userId)

	if err := row.Scan(&teacherId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf(`insert into %s (role_id, user_id) values ($1, $2)`, usersRolesTable)
	_, err = tx.Exec(query, teacherRoleId, userId)

	if err != nil {
		return 0, err
	}

	return teacherId, tx.Commit()
}

func (r repository) DeleteTeacher(teacherId int) error {
	return nil
}

func (r repository) GetTeacher(teacherId int) (dtos.User, error) {
	var teacher dtos.User
	return teacher, nil
}

func (r repository) GetTeachers() ([]dtos.User, error) {
	return nil, nil
}
