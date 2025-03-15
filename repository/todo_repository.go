package repository

import (
	"context"
	"database/sql"
	"eisenhower-todo-api/model/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, tx *sql.Tx, Todo domain.Todo) domain.Todo
	Patch(ctx context.Context, tx *sql.Tx, Todo domain.Todo) domain.Todo
	Delete(ctx context.Context, tx *sql.Tx, todoId int)
	FindById(ctx context.Context, tx *sql.Tx, todoId int) domain.Todo
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo
}
