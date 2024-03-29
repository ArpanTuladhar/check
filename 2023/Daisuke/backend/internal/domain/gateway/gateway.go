package gateway

//go:generate moq -out mock_gateway.go . TodoCommandsGateway Binder Transactor

import (
	"context"

	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/model/todo"
)

type Binder interface {
	Bind(context.Context) context.Context
}

type Transactor interface {
	Transaction(context.Context, func(context.Context) error) error
}

type TodoCommandsGateway interface {
	CreateTodo(
		ctx context.Context,
		newTodo *todo.NewTodo,
	) (*todo.Todo, error)
}
