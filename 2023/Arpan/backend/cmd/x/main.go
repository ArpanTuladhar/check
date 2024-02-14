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

	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/domain/service"
	h "github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/handler"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/infrastructure/todo"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/middleware"
	"github.com/88labs/andpad-engineer-training/2023/Arpan/backend/internal/utils/config"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	addr := fmt.Sprintf(":%d", conf.ServerPort)
	listener, err2 := net.Listen("tcp", addr)
	if err2 != nil {
		log.Fatalf("Error starting listener: %v", err2)
	}

	ownerConn, cleanup, err := todo.NewOwnerSQLHandler(conf)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	binder := todo.NewConnectionBinder(ownerConn)
	transactor := todo.NewTransactor(ownerConn)

	todoWriter := todo.NewTodoWriter()
	todoCreator := service.NewTodoCreator(binder, transactor, todoWriter)

	middle := middleware.NewMiddleware()
	router := h.NewHTTPServer(conf, middle, todoCreator)

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

	fmt.Println("started todo server")
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
