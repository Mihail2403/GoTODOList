package service

import (
	"go_todo_list/entity"
	"go_todo_list/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(userId int, list entity.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
