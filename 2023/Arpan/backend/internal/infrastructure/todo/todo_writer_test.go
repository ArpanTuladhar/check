package todo

import (
	"context"
	"testing"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/gateway"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/session"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/model/todo"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/utils/config"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestTodoWriter_CreateTodo_Success(t *testing.T) {
	// Load configuration
	conf, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}
	mockGateway := &gateway.TodoCommandsGatewayMock{
		CreateFunc: func(ctx context.Context, newTodo *todo.NewTodo) (*todo.Todo, error) {
			return &todo.Todo{
				ID:     todo.TodoID(uuid.NewString()),
				Text:   newTodo.Text,
				UserID: 123456,
			}, nil
		},
	}
	todoWriter := NewTodoWriter(mockGateway)
	sess := &session.Session{UserId: 123456}
	ctx := session.StoreSession(context.Background(), sess)
	// Create a mock GORM DB instance using the loaded configuration
	dsn := config.ConstructDSN(conf)
	conn, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}))
	if err != nil {
		t.Fatalf("failed to open mock DB connection: %v", err)
	}
	ctx = WithTodoDB(ctx, conn)
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
