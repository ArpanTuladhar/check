package main

import (
	"context"
	"fmt"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/domain/service"
	h "github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/handler"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/infrastructure/todo"
	"github.com/88labs/andpad-engineer-training/2023/Daisuke/backend/internal/middleware"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	addr := fmt.Sprintf(":%s", port)
	listener, err2 := net.Listen("tcp", addr)
	if err2 != nil {
		panic(err2)
	}

	todoWriter := todo.NewTodoWriter()
	todoCreator := service.NewTodoCreator(todoWriter)

	middle := middleware.NewMiddleware()
	router := h.NewHTTPServer(middle, todoCreator)

	ch := make(chan error)
	go func() {
		srv := &http.Server{
			Handler:           router,
			ReadTimeout:       15 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       30 * time.Second,
		}
		ch <- srv.Serve(listener)
	}()

	fmt.Println("started todo server")
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-sigCh:
			_ = listener.Close()
		case err := <-ch:
			_ = listener.Close()
			fmt.Println("error!!", err.Error())
		}
		cancel()
	}()
	<-ctx.Done()

}
