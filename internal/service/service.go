package service

import (
	"github.com/Lunovoy/todo-api/internal/model"
	"github.com/Lunovoy/todo-api/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Autorization),
		TodoList:      NewTodoListService(repo.TodoList),
	}
}
