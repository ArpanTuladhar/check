package service

import (
	"context"
	"fmt"

	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/usecase"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/usecase/output"
)

type todoCreator struct {
	todoCommandsGateway gateway.TodoCommandsGateway
}

func NewTodoCreator(todoCommandsGateway gateway.TodoCommandsGateway) usecase.TodoCreator {
	return &todoCreator{todoCommandsGateway}
}

func (t todoCreator) CreateTodo(ctx context.Context, in *input.TodoCreator) (*output.TodoCreator, error) {
	todo, err := t.todoCommandsGateway.Create(ctx, &todo.NewTodo{Text: "text"})
	if err != nil {
		return nil, fmt.Errorf("create todo failed: %w", err)
	}

	return &output.TodoCreator{ID: todo.ID, Text: todo.Text, UserID: todo.UserID}, nil
}
