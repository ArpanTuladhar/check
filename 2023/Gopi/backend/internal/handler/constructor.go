package handler

import (
	"net/http"

	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph"
	generated "github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph/generated"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/usecase"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
)

func NewHTTPServer(todoCreator usecase.TodoCreator) http.Handler {
	router := chi.NewRouter()
	router.Route("/graphql", func(r chi.Router) {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(todoCreator)))
		r.Handle("/", srv)
	})

	return router
}
