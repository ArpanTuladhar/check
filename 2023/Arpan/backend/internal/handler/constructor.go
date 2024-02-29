package handler

import (
	"net/http"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/utils/config"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph"
	generated "github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph/generated"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/middleware"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/usecase"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
)

func NewHTTPServer(c *config.Config, middle middleware.Middleware, todoCreator usecase.TodoCreator) http.Handler {
	router := chi.NewRouter()
	router.Route("/graphql", func(r chi.Router) {
		r.Use(
			middle.WithAuth(),
		)
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(todoCreator)))
		r.Handle("/", srv)
	})

	if c.Env != config.EnvProduction {
		router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	}

	return router
}
