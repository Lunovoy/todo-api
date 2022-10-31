package service

import (
	"github.com/Lunovoy/todo-api/internal/model"
	"github.com/Lunovoy/todo-api/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list model.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
