package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable     = "users"
	todoListTable = "lists"
	todoItemTable = "items"
	userListTable = "user_list"
	listItemTabel = "list_item"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
