package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skill-management-api/database"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type Skill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

func getSkill(ctx *gin.Context) {
	fmt.Println("Entering getSkill handler")
	skills := []Skill{}

	rows, err := database.DB.Query("SELECT key, name, description, logo, tags FROM skill")
	if err != nil {
		log.Fatal("can't query all skills", err)
	}

	for rows.Next() {
		var key, name, description, logo string
		var tags pq.StringArray

		err := rows.Scan(&key, &name, &description, &logo, &tags)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		fmt.Println(key, name, description, logo, tags)
		skills = append(skills, Skill{key, name, description, logo, tags})
	}

	fmt.Println("query all skills success")
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": skills})
}

func getSkillByKey(ctx *gin.Context) {
	fmt.Println("Entering getSkillByID handler")

	key := ctx.Param("key")
	q := "SELECT key, name, description, logo, tags FROM skill where key=$1"

	row := database.DB.QueryRow(q, key)
	var name, description, logo string
	var tags pq.StringArray

	err := row.Scan(&key, &name, &description, &logo, &tags)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	fmt.Println("one row", key, name, description, logo, tags)
	ctx.JSON(http.StatusOK, gin.H{"data": Skill{key, name, description, logo, tags}, "status": "success"})

}

func postSkill(ctx *gin.Context) {
	fmt.Println("Entering postSkill handler")
	var skill Skill

	if err := ctx.BindJSON(&skill); err != nil {
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key"
	err := database.DB.QueryRow(query, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags)).Scan(&skill.Key)
	if err != nil {
		fmt.Println("Error inserting new skill:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Skill created with Key:", skill.Key)
	// ctx.JSON(http.StatusCreated, skill)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": skill})
}

func putSkillByKey(ctx *gin.Context) {
	fmt.Println("Entering putSkillByKey handler")
	key := ctx.Param("key")
	var skill Skill

	if err := ctx.BindJSON(&skill); err != nil {
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE skill SET name=$2, description=$3, logo=$4, tags=$5 WHERE key=$1 RETURNING key;"
	if _, err := database.DB.Exec(query, key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags)); err != nil {
		fmt.Println("Error executing update:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Skill update with Key:", skill.Key)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": skill})
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	database.ConnectDB()
	defer database.DB.Close()
	database.CreateTable()

	r := gin.Default()
	r.GET("/api/v1/skills", getSkill)
	r.GET("/api/v1/skills/:key", getSkillByKey)
	r.POST("/api/v1/skills", postSkill)
	r.PUT("/api/v1/skills/:key", putSkillByKey)
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
