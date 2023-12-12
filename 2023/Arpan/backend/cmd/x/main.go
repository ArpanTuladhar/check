package main

import (
	"log"
	"net/http"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/config"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/service"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph"
	generated "github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	appConfig := config.LoadAppConfig()
	port := appConfig.Port

	todoCreator := service.NewTodoCreator()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(todoCreator)))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
