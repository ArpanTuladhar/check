package usecase

import (
	"context"
	
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/usecase/output"
)

type TodoCreator interface {
	CreateTodo(ctx context.Context, in *input.TodoCreator) (*output.TodoCreator, error)
}