package repository

import (
	"BilimSearch/dtos"
	"BilimSearch/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	log.Printf("started transaction")

	lessons, err := r.GetTeachersLessons(tx, teacherId)

	if err != nil {
		return err
	}
	log.Printf("get lessons id")
	for i := 0; i < len(lessons); i++ {
		log.Print(lessons[i])
	}

	err = r.UpdateLessonsTeacher(tx, lessons)

	if err != nil {
		return err
	}
	log.Printf("updated all lessons")

	userId, err := r.GetIdByRole(tx, teacherId, teachersTable)
	if err != nil {
		return err
	}
	log.Printf("get userId")

	query := fmt.Sprintf("delete from %s where id = $1", teachersTable)
	_, err = tx.Exec(query, teacherId)

	if err != nil {
		return err
	}
	log.Printf("deleted!")
	query = fmt.Sprintf("delete from %s where user_id = $1 and role_id = $2", usersRolesTable)
	_, err = tx.Exec(query, userId, teacherRoleId)

	if err != nil {
		return err
	}
	log.Printf("deleted roles!")

	return tx.Commit()
}

func (r repository) GetTeacher(teacherId int) (dtos.User, error) {
	var teacher dtos.User
	return teacher, nil
}

func (r repository) GetTeachers() ([]dtos.User, error) {
	return nil, nil
}

func (r repository) GetTeachersLessons(tx *sqlx.Tx, teacherId int) ([]int, error) {
	var lessons []int
	query := fmt.Sprintf("select id from %s where teacher_id = $1", lessonsTable)
	rows, err := tx.Queryx(query, teacherId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for rows.Next() {
		var lessonId int
		err := rows.Scan(&lessonId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		lessons = append(lessons, lessonId)
	}
	return lessons, nil
}

func (r repository) UpdateLessonsTeacher(tx *sqlx.Tx, lessons []int) error {
	query := fmt.Sprintf("update %s set teacher_id = 0 where id = any($1)", lessonsTable)
	_, err := tx.Exec(query, pq.Array(lessons))
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
