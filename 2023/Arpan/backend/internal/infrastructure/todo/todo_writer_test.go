package todo

import (
	"context"
	"errors"
	"testing"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/session"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
	"github.com/stretchr/testify/assert"
)

type TodoWriter struct{}

func NewTodoWriterInstance() *TodoWriter {
	return &TodoWriter{}
}

func (tw *TodoWriter) CreateTodo(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
	if newTodo.Text == "" {
		return nil, errors.New("todo text cannot be empty")
	}
	sess, err := session.ExtractSession(ctx)
	if err != nil {
		return nil, err
	}
	createdTodo := &todo.Todo{
		ID:     "test_todo_id",
		Text:   newTodo.Text,
		UserID: sess.UserId,
	}
	return createdTodo, nil
}
func TestTodoWriter_CreateTodo_Success(t *testing.T) {
	sess := &session.Session{UserId: 123456}
	ctx := session.WithSession(context.Background(), sess)
	todoWriter := NewTodoWriterInstance()
	newTodo := &todo.NewTodo{Text: "todo_text"}
	gotTodo, gotErr := todoWriter.CreateTodo(ctx, newTodo)
	expectedTodo := &todo.Todo{
		ID:     "test_todo_id",
		Text:   "todo_text",
		UserID: 123456,
	}
	assert.Equal(t, expectedTodo, gotTodo)
	assert.NoError(t, gotErr)
}
