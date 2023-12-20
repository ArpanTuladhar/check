package main


import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"

	"github.com/julienschmidt/httprouter"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/domain/service"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/handler/graph"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/infrastructure/todo"
	generated "github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/handler/graph/generated"
)

const defaultPort = "8080"

func newRouter() *httprouter.Router {
	router := httprouter.New()

	// GraphQL playground route
	router.GET("/", playgroundHandler())
	
	// GraphQL query route
	router.POST("/query", graphqlHandler())

	return router
}

func playgroundHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(w, r)
	}
}

func graphqlHandler() httprouter.Handle {
	todoWriter := todo.NewTodoWriter()
	todoCreator := service.NewTodoCreator(todoWriter)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graph.New(todoCreator)))

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		srv.ServeHTTP(w, r)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := newRouter()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on http://localhost:%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	sig := <-signalCh
	log.Printf("Received signal: %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
		return
	}

	log.Println("Server shutdown gracefully")
}
