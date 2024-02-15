package integration

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/testutil"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/utils/config"
	"gorm.io/gorm"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/service"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/infrastructure/todo"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/middleware"
)

func DoGraphQLRequest(
	ctx context.Context,
	buffer *bytes.Buffer,
	gormDB *gorm.DB,
	dbName string,
) *httptest.ResponseRecorder {
	router := initHttpServerForIntegrationTest(
		ctx,
		gormDB,
		dbName,
	)

	recorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/graphql", buffer)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)
	return recorder
}

func initHttpServerForIntegrationTest(ctx context.Context, gormDB *gorm.DB, dbName string) http.Handler {
	middle := middleware.NewMiddleware()

	todoDBConn := &todo.TodoConn{GormDB: gormDB}

	e := testutil.LoadEnv()

	conf := &config.Config{
		ServerPort: 50051,
		DBHost:     e.DBHost,
		DBPort:     e.DBPort,
		DBUser:     e.DBUser,
		DBPass:     e.DBPass,
		DBName:     dbName,
	}

	binder := todo.NewConnectionBinder(todoDBConn)
	transactor := todo.NewTransactor(todoDBConn)

	todoWriter := todo.NewTodoWriter()
	todoCreator := service.NewTodoCreator(binder, transactor, todoWriter)
	handler := handler.NewHTTPServer(conf, middle, todoCreator)
	return handler
}
