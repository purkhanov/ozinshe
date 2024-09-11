package repository

import (
	"fmt"
	"ozinshe/schemas"

	"strings"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAllUsers() ([]schemas.User, error) {
	var users []schemas.User

	query := fmt.Sprintf(
		`
		SELECT 
			u.id, 
			u.first_name, 
			u.email, 
			u.is_admin, 
			u.phone_number, 
			u.year_of_birth, 
			u.created_at 
		FROM %s u;
		`,
		usersTable,
	)

	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPostgres) GetUser(userId int) (schemas.User, error) {
	var user schemas.User

	query := fmt.Sprintf(
		`
		SELECT 
			id, 
			first_name, 
			email, 
			is_admin, 
			phone_number, 
			year_of_birth, 
			created_at 
		FROM %s 
		WHERE id = $1;
		`,
		usersTable,
	)

	if err := r.db.Get(&user, query, userId); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserPostgres) UpdateUser(userId int, input map[string]any) error {
	if len(input) == 0 {
		return nil
	}

	setKeysVal := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	for key, val := range input {
		sv := fmt.Sprintf("%s=$%d", key, argId)
		setKeysVal = append(setKeysVal, sv)
		args = append(args, val)
		argId++
	}

	setQuery := strings.Join(setKeysVal, ", ")
	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = $%d",
		usersTable, setQuery, argId,
	)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *UserPostgres) DeleteUser(userId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1",
		usersTable,
	)
	_, err := r.db.Exec(query, userId)
	return err
}
