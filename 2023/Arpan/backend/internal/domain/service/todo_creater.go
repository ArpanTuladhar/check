package service

import (
	"context"
	"fmt"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/output"
)

type todoCreator struct {
	todoCommandsGateway gateway.TodoCommandsGateway
}

func NewTodoCreator(todoCommandsGateway gateway.TodoCommandsGateway) usecase.TodoCreator {
	return &todoCreator{todoCommandsGateway}
}

func (t todoCreator) CreateTodo(ctx context.Context, in *input.TodoCreator) (*output.TodoCreator, error) {
	todo, err := t.todoCommandsGateway.Create(ctx, &todo.NewTodo{Text: in.Text})
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return &output.TodoCreator{ID: todo.ID, Text: todo.Text, UserID: todo.UserID}, nil
}
