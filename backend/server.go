package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	// r.GET("/api/v1/todos", getTodo)
	// r.GET("/api/v1/todos/:id", getTodoByID)
	// r.POST("/api/v1/todos", postTodo)
	// r.PUT("/api/v1/todos/:id", putTodoByID)
	// r.DELETE("/api/v1/todos/:id", deleteTodoByID)
	// r.PATCH("/api/v1/todos/:id/actions/status", patchTodoStatusByID)
	// r.PATCH("/api/v1/todos/:id/actions/title", patchTodoTitleByID)

	port := os.Getenv("HOST")
	// if port == "" {
	// 	fmt.Println("Why Port Is String?!!")
	// 	port = "8080" // default port if not specified
	// }

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	serverErrors := make(chan error, 1)

	// Start the service listening for requests
	go func() {
		log.Printf("Listening on port %s", port)
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		log.Println("Received shutdown signal, gracefully shutting down...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("Graceful shutdown failed: %v", err)
		}

	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}

	log.Println("Server stopped")
}
