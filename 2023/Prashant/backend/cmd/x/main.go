package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	graph "github.com/88labs/andpad-engineer-training/2023/Prashant/backend/internal"
	generated "github.com/88labs/andpad-engineer-training/2023/Prashant/backend/internal/handler/graph/generated"
)

func waitForInterrupt(ctx context.Context, server *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Blocked until a signal is received or the context is canceled.
	select {
	case sig := <-c:
		fmt.Printf("Received signal: %v\n", sig)
	case <-ctx.Done():
		fmt.Println("Context was cancelled")
		return
	}

	// context with a timeout for graceful shutdown.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// shut down the server gracefully.
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Error during server shutdown: %v\n", err)
		return
	}
	fmt.Println("Server gracefully shut down")
}

func main() {
	// context that listens for interrupt and terminate signals.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//  default port and read from the environment variable.
	const defaultPort = "8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// GraphQL server with generated schema and resolver.
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Start the HTTP server in a goroutine.
	server := &http.Server{
		Addr:    ":" + port,
		Handler: srv,
	}

	go func() {
		fmt.Printf("Server is starting on http://localhost:%s\n", port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Start http.server failed: %v\n", err)
			cancel()
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server.
	waitForInterrupt(ctx, server)
}
