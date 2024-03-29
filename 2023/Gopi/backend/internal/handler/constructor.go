package handler

import (
	"net/http"

	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph"
	generated "github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph/generated"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/middleware"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
)

func NewHTTPServer(middle middleware.Middleware, todoCreator usecase.TodoCreator) http.Handler {
	router := chi.NewRouter()
	router.Route("/graphql", func(r chi.Router) {
		r.Use(
			middle.WithAuth(),
		)
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(todoCreator)))
		r.Handle("/", srv)
	})

	return router
}
