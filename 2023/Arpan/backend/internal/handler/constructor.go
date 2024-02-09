package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph"
	generated "github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph/generated"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/middleware"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase"
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
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	return router
}
