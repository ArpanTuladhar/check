package todo

import (
	"context"
	"errors"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/session"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
)

type todoWriter struct {
}

func NewTodoWriter() gateway.TodoCommandsGateway {
	return &todoWriter{}
}

func (t todoWriter) Create(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
	s, err := session.ExtractSession(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch a session")
	}
	return &todo.Todo{ID: "todo_id_1", Text: "todo_text_1", UserID: s.UserId}, nil
}
