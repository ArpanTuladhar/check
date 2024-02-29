package service

import (
	"context"
	"errors"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/output"
)

type todoCreator struct {
	dbConnBinder        gateway.Binder
	todoCommandsGateway gateway.TodoCommandsGateway
}

func NewTodoCreator(dbConnBinder gateway.Binder, todoCommandsGateway gateway.TodoCommandsGateway) usecase.TodoCreator {
	return &todoCreator{dbConnBinder, todoCommandsGateway}
}

func (t todoCreator) CreateTodo(ctx context.Context, in *input.TodoCreator) (*output.TodoCreator, error) {
	ctx = t.dbConnBinder.Bind(ctx)

	todo, err := t.todoCommandsGateway.Create(ctx, &todo.NewTodo{Text: in.Text})
	if err != nil {
		return nil, errors.New("error")
	}

	return &output.TodoCreator{ID: todo.ID, Text: todo.Text, UserID: todo.UserID}, nil
}
