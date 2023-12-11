package main

import (
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/domain/service"
	"github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph"
	generated "github.com/88labs/andpad-engineer-training/2023/Gopi/backend/internal/handler/graph/generated"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	todoCreator := service.NewTodoCreator()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(todoCreator)))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
