package todo

import (
	"context"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
)

type todoWriter struct {
}

func NewTodoWriter() gateway.TodoCommandsGateway {
	return &todoWriter{}
}

func (t todoWriter) CreateTodo(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
	return &todo.Todo{ID: "id", Text: "test"}, nil
}
