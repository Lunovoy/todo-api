package repository

import (
	"github.com/Lunovoy/todo-api/internal/model"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(driver model.User) (int, error)
	GetUser(login, password string) (model.User, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
}

type TodoItem interface {
}

type Repository struct {
	Autorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
		TodoList:     NewTodoListPostgres(db),
	}
}
