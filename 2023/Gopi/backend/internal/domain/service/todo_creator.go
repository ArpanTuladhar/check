package service

import (
	"context"
	"errors"

	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase/output"
)

type todoCreator struct {
	todoCommandsGateway gateway.TodoCommandsGateway
}

func NewTodoCreator(todoCommandsGateway gateway.TodoCommandsGateway) usecase.TodoCreator {
	return &todoCreator{todoCommandsGateway}
}

func (t todoCreator) CreateTodo(ctx context.Context, in *input.TodoCreator) (*output.TodoCreator, error) {
	todo, err := t.todoCommandsGateway.CreateTodo(ctx, &todo.NewTodo{Text: "text"})
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return &output.TodoCreator{ID: todo.ID, Text: todo.Text, UserID: todo.UserID}, nil
}
