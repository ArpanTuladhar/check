package handler

import (
	"net/http"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph"
	generated "github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph/generated"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
)

func NewHTTPServer(todoCreator usecase.TodoCreator) http.Handler {
	router := chi.NewRouter()
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(todoCreator))))
	return router
}
