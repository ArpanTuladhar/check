package gateway

import (
	"context"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
)

type TodoCommandsGateway interface {
	Create(
		ctx context.Context,
		newTodo *todo.NewTodo,
	) (*todo.Todo, error)
}
