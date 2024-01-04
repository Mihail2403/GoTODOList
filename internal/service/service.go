package service

import (
	"go_todo_list/entity"
	"go_todo_list/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Update(userId, listId int, input entity.UpdateTodoListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(userId, listId int, item entity.TodoItem) (int, error)
	GetAllByList(userId, listId int) ([]entity.TodoItem, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewItemService(repos.TodoItem, repos.TodoList),
	}
}
