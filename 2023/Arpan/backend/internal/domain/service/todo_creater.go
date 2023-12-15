package service

import (
	"context"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/output"
)

type todoCreator struct {
}

func NewTodoCreator() usecase.TodoCreator {
	return &todoCreator{}
}

func (t todoCreator) CreateTodo(ctx context.Context, in *input.TodoCreator) (*output.TodoCreator, error) {
	//TODO implement me
	panic("implement me")
}
