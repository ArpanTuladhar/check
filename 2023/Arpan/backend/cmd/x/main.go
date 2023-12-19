package main

import (
	"log"
	"net/http"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/config"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/service"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph"
	generated "github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler/graph/generated"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/infrastructure/datastore"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/infrastructure/todo"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	// Connect to the database
	db, err := datastore.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	appConfig := config.LoadAppConfig()
	port := appConfig.Port

	todoWriter := todo.NewTodoWriter()
	todoCreator := service.NewTodoCreator(todoWriter)
	//FIXME: create a constructor.go file in the handler module
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(todoCreator)))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Connected to the database. Now, connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
