package graph

import (
	generated "github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph/generated"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	todoCreator usecase.TodoCreator
}

func New(todoCreator usecase.TodoCreator) generated.Config {
	return generated.Config{Resolvers: &Resolver{todoCreator: todoCreator}}
}
