package repository

import (
	"BilimSearch/models"
	"database/sql"
	"fmt"
	"strings"
)

func (r *repository) GetRoles(id int) ([]string, error) {
	var roles []string
	query := fmt.Sprintf("select r.role_name from %s r join %s c on c.role_id = r.id where c.user_id = $1", rolesTable, usersRolesTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return roles, err
	}
	for rows.Next() {
		var role string
		err := rows.Scan(&role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, err
}

func (r *repository) GetId(role string, user_id int) (int, error) {
	var id int
	table := "t_" + strings.ToLower(role) + "s"

	query := fmt.Sprintf("select id from %s where user_id = $1", table)
	err := r.db.Get(&id, query, user_id)
	return id, err
}  

func (r *repository) GetUser(username string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id, username, password from %s where username = $1", usersTable)
	err := r.db.Get(&user, query, username)
	return user, err
}

func (r *repository) CreateUser(user models.User, tx *sql.Tx) (int, error) {
	var user_id int
	query := fmt.Sprintf(fmt.Sprintf("insert into %s (username, password, name, surname, email) values ($1, $2, $3, $4, $5) returning id", usersTable))
	row := tx.QueryRow(query, user.Username, user.Password, user.Name, user.Surname, user.Email)
	if err := row.Scan(&user_id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return user_id, nil
}