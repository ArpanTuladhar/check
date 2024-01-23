package usecase

//go:generate moq -out mock_usecase.go . TodoCreator

import (
	"context"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/input"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase/output"
)

type TodoCreator interface {
	CreateTodo(ctx context.Context, in *input.TodoCreator) (*output.TodoCreator, error)
}
