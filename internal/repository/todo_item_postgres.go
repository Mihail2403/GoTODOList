package repository

import (
	"fmt"
	"go_todo_list/entity"

	"github.com/jmoiron/sqlx"
)

type TodoItemPostgresRepo struct {
	db *sqlx.DB
}

func NewTodoItemPostgresRepo(db *sqlx.DB) *TodoItemPostgresRepo {
	return &TodoItemPostgresRepo{
		db: db,
	}
}

func (r *TodoItemPostgresRepo) Create(listId int, item entity.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id`, todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {

		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf(`INSERT INTO %s (list_id, item_id) VALUES ($1, $2)`, listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()

		return 0, err
	}
	return itemId, tx.Commit()
}

func (r *TodoItemPostgresRepo) GetAllByList(userId, listId int) ([]entity.TodoItem, error) {
	var items []entity.TodoItem
	query := fmt.Sprintf(`
	SELECT ti.id, ti.title, ti.description, ti.done 
	FROM %s ti
		INNER JOIN %s li ON li.item_id = ti.id
		INNER JOIN %s ul ON ul.list_id = li.list_id
	WHERE li.list_id = $1 AND ul.user_id = $2
	`, todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.Select(&items, query, listId, userId)
	if err != nil {
		return nil, err
	}

	return items, nil
}
