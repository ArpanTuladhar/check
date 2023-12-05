package gateway

import (
	"context"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/model/todo"
)

type TodoCommandsGateway interface {
	CreateTodo(
		ctx context.Context,
		newTodo *todo.NewTodo,
	) (*todo.Todo, error)
}
