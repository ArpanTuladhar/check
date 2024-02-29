package todo

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/session"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/testutil"
)

func TestTodoWriter_CreateTodo_Success(t *testing.T) {

	db, _ := testutil.InitDB(t)

	ctx := WithTodoDB(context.Background(), db)
	todoWriter := NewTodoWriter()

	sess := &session.Session{UserId: 123456}
	ctx = session.StoreSession(ctx, sess)

	newTodo := &todo.NewTodo{Text: "todo_text"}

	gotTodo, gotErr := todoWriter.Create(ctx, newTodo)

	expectedTodo := &todo.Todo{
		ID:     todo.TodoID(uuid.NewString()),
		Text:   "todo_text",
		UserID: 123456,
	}

	assert.Equal(t, expectedTodo.Text, gotTodo.Text)
	assert.Equal(t, expectedTodo.UserID, gotTodo.UserID)
	assert.NoError(t, gotErr)
}
