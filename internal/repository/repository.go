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
}

type TodoItem interface {
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
	}
}
