package repository

import (
	"BilimSearch/models"
	"database/sql"
	"fmt"
)

func (r *repository) GetUser(username string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id, role_id from %s where username = $1", usersTable)
	err := r.db.Get(user, query, username)
	return user, err
}

func (r *repository) CreateUser(username, password string, tx *sql.Tx) (int, error) {
	var user_id int
	query := fmt.Sprintf(fmt.Sprintf("insert into %s (username, password) values ($1, $2) returning id", usersTable))
	row := tx.QueryRow(query, username, password)
	if err := row.Scan(&user_id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return user_id, nil
}