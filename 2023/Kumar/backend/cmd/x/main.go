package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/domain/service"
	h "github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/handler"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/infrastructure/todo"
	"github.com/88labs/andpad-engineer-training/2023/Kumar/backend/internal/middleware"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	addr := fmt.Sprintf(":%s", port)

	todoWriter := todo.NewTodoWriter()
	todoCreator := service.NewTodoCreator(todoWriter)

	middle := middleware.NewMiddleware()
	router := h.NewHTTPServer(middle, todoCreator)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		fmt.Println("Starting todo server on", addr)
		if err := srv.Serve(listener); err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
		cancel()
	}()

	select {
	case sig := <-signalCh:
		fmt.Println("Received signal:", sig)
	case <-ctx.Done():
	}

	// Shutdown the server gracefully
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Println("Server shutdown error:", err)
	} else {
		fmt.Println("Server shutdown gracefully")
	}
}
