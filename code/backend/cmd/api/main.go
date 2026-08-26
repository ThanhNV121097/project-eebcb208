package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-eebcb208/backend/internal/migrations"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type app struct {
	db *pgxpool.Pool
}

type errorBody struct {
	Error errorDetail `json:"error"`
}

type errorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type greetingBody struct {
	Text string `json:"text"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("create db pool: %v", err)
	}
	defer db.Close()

	if err := applyMigrations(ctx, db); err != nil {
		log.Fatalf("apply migrations: %v", err)
	}
	if err := db.Ping(ctx); err != nil {
		log.Fatalf("ping db: %v", err)
	}

	a := app{db: db}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", a.healthz)
	mux.HandleFunc("GET /v1/greeting", a.greeting)

	server := &http.Server{
		Addr:              ":" + port(),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %v", err)
	}
}

func port() string {
	if value := os.Getenv("PORT"); value != "" {
		return value
	}
	if value := os.Getenv("APP_PORT"); value != "" {
		return value
	}
	return "8080"
}

func applyMigrations(ctx context.Context, db *pgxpool.Pool) error {
	entries, err := migrations.Files.ReadDir(".")
	if err != nil {
		return err
	}

	var names []string
	for _, entry := range entries {
		name := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(name, ".up.sql") {
			names = append(names, name)
		}
	}
	sort.Strings(names)

	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (
		version text PRIMARY KEY,
		applied_at timestamptz NOT NULL DEFAULT now()
	)`)
	if err != nil {
		return err
	}

	for _, name := range names {
		if err := applyMigration(ctx, db, name); err != nil {
			return err
		}
	}
	return nil
}

func applyMigration(ctx context.Context, db *pgxpool.Pool, name string) error {
	version := strings.TrimSuffix(name, ".up.sql")
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var exists bool
	if err := tx.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)", version).Scan(&exists); err != nil {
		return err
	}
	if exists {
		return tx.Commit(ctx)
	}

	sqlBytes, err := migrations.Files.ReadFile(name)
	if err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	if _, err := tx.Exec(ctx, "INSERT INTO schema_migrations (version) VALUES ($1)", version); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (a app) healthz(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := a.db.Ping(ctx); err != nil {
		writeError(w, http.StatusServiceUnavailable, "unhealthy", "Service is not healthy.")
		return
	}

	var ok int
	if err := a.db.QueryRow(ctx, "SELECT 1").Scan(&ok); err != nil || ok != 1 {
		writeError(w, http.StatusServiceUnavailable, "unhealthy", "Service is not healthy.")
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok"))
}

func (a app) greeting(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var text string
	err := a.db.QueryRow(ctx, "SELECT text FROM greetings WHERE id = $1", 1).Scan(&text)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusNotFound, "greeting_not_found", "Greeting was not found.")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal_error", "Something went wrong.")
		return
	}

	writeJSON(w, http.StatusOK, greetingBody{Text: text})
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, errorBody{Error: errorDetail{Code: code, Message: message}})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
