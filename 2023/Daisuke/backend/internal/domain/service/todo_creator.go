package service

import (
	"context"
	"errors"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/utils/config"

	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/usecase"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/usecase/output"
)

type todoCreator struct {
	conf                *config.Config
	dbConnBinder        gateway.Binder
	transactor          gateway.Transactor
	todoCommandsGateway gateway.TodoCommandsGateway
}

func NewTodoCreator(conf *config.Config, dbConnBinder gateway.Binder, transactor gateway.Transactor, todoCommandsGateway gateway.TodoCommandsGateway) usecase.TodoCreator {
	return &todoCreator{conf, dbConnBinder, transactor, todoCommandsGateway}
}

func (t todoCreator) CreateTodo(ctx context.Context, in *input.TodoCreator) (*output.TodoCreator, error) {
	ctx = t.dbConnBinder.Bind(ctx)

	todo, err := t.todoCommandsGateway.CreateTodo(ctx, &todo.NewTodo{Text: in.Text})
	if err != nil {
		return nil, errors.New("error")
	}

	return &output.TodoCreator{ID: todo.ID, Text: todo.Text, UserID: todo.UserID}, nil
}
