package todo

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/session"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
)

type todoWriter struct {
}

func NewTodoWriter() gateway.TodoCommandsGateway {
	return &todoWriter{}
}

func (t todoWriter) CreateTodo(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
	s, err := session.ExtractSession(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch a session")
	}
	tx, err := ExtractTodoDB(ctx)
	if err != nil {
		return nil, err
	}

	createdTodo := &todo.Todo{
		ID:     todo.TodoID(uuid.NewString()),
		Text:   newTodo.Text,
		UserID: s.UserId,
	}

	db := tx.WithContext(ctx)
	if err := db.
		Create(&createdTodo).
		Take(&createdTodo).Error; err != nil {
		return nil, err
	}

	return createdTodo, nil
}
