package integration

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/domain/service"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/infrastructure/todo"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/middleware"
)

func DoGraphQLRequest(
	buffer *bytes.Buffer,
) *httptest.ResponseRecorder {

	router := initHttpServerForIntegrationTest(
		context.Background(),
	)

	recorder := httptest.NewRecorder()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/graphql", buffer)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(recorder, req)
	return recorder
}

func initHttpServerForIntegrationTest(ctx context.Context) http.Handler {
	middle := middleware.NewMiddleware()
	todoWriter := todo.NewTodoWriter()
	todoCreator := service.NewTodoCreator(todoWriter)
	handler := handler.NewHTTPServer(middle, todoCreator)
	return handler
}
