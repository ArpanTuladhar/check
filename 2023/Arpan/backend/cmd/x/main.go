package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/config"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/service"
	h "github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/infrastructure/todo"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/middleware"
)

func main() {
	// Connect to the database
	configObj, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Error loading app configuration: %v", err)
	}

	db, err := configObj.Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	port := configObj.App.Port

	todoWriter := todo.NewTodoWriter()
	todoCreator := service.NewTodoCreator(todoWriter)

	middle := middleware.NewMiddleware()
	router := h.NewHTTPServer(middle, todoCreator)

	addr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}

	ch := make(chan error)
	go func() {
		srv := &http.Server{
			Handler:           router,
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       30 * time.Second,
		}

		if err := srv.Serve(listener); err != nil {
			ch <- err
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-sigCh:
			log.Println("Received interrupt signal. Shutting down...")
			_ = listener.Close()
		case err := <-ch:
			log.Printf("Server error: %v", err)
			_ = listener.Close()
		}
		cancel()
	}()

	<-ctx.Done()
	log.Println("Server gracefully shut down")
}
