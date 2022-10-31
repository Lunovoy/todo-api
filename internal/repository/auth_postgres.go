package repository

import (
	"fmt"

	"github.com/Lunovoy/todo-api/internal/model"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, login, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, login, password)

	return user, err
}
