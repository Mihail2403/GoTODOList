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

func (s *TodoListService) GetAll(userId int) ([]entity.TodoList, error) {
	return s.repo.GetAll(userId)
}
func (s *TodoListService) GetById(userId, listId int) (entity.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
func (s *TodoListService) Update(userId, listId int, input entity.UpdateTodoListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
