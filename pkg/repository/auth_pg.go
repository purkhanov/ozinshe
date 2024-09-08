package repository

import (
	"fmt"
	"ozinshe/schemas"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user schemas.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		`
		INSERT INTO %s (email, password_hash)
		VALUES ($1, $2) 
		RETURNING id
		`,
		usersTable,
	)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (schemas.User, error) {
	var user schemas.User

	query := fmt.Sprintf(
		`
		SELECT id, is_admin  
		FROM %s 
		WHERE email=$1 AND password_hash=$2
		`,
		usersTable,
	)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
