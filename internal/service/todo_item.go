package service

import (
	"go_todo_list/entity"
	"go_todo_list/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewItemService(repo repository.TodoItem, list_repo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: list_repo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item entity.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAllByList(userId, listId int) ([]entity.TodoItem, error) {
	return s.repo.GetAllByList(userId, listId)
}
