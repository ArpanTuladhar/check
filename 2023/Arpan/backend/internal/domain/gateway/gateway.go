package gateway

//go:generate moq -out mock_gateway.go . TodoCommandsGateway

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
