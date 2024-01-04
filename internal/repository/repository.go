package repository

import (
	"go_todo_list/entity"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username string, password string) (entity.User, error)
}

type TodoList interface {
	Create(userId int, list entity.TodoList) (int, error)
	GetAll(userId int) ([]entity.TodoList, error)
	GetById(userId, listId int) (entity.TodoList, error)
	Update(userId, listId int, input entity.UpdateTodoListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(listId int, item entity.TodoItem) (int, error)
	GetAllByList(userId, listId int) ([]entity.TodoItem, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgresRepo(db),
		TodoItem:      NewTodoItemPostgresRepo(db),
	}
}
