package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/romietis/lunar-plan-advisor/v4/advisor"
	"github.com/romietis/lunar-plan-advisor/v4/internal/endpoints"
)

func main() {
	// Load default plans from disk.
	data, err := os.ReadFile("plans.json")
	if err != nil {
		log.Fatal(err)
	}
	var defaults advisor.PlansConfig
	if err := json.Unmarshal(data, &defaults); err != nil {
		log.Fatal(err)
	}

	// Parse HTML templates.
	tmpl := template.Must(template.ParseGlob("assets/templates/*"))

	// Build the handler with templates plus default plans.
	handlers := &endpoints.Handler{
		Defaults:  defaults,
		Templates: tmpl,
	}

	// Register routes: static assets, then app endpoints.
	mux := http.NewServeMux()
	mux.Handle("GET /css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css"))))
	mux.Handle("GET /js/", http.StripPrefix("/js/", http.FileServer(http.Dir("assets/js"))))
	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /plans", handlers.GetPlans)
	mux.HandleFunc("POST /plans/best", handlers.PostBestPlans)

	// Configure the HTTP server with timeouts.
	// Read/Write/ReadHeader timeouts defend against slow-client attacks (slowloris, slow body, slow read).
	// IdleTimeout bounds keep-alive idle connections.
	addr := ":" + port()
	srv := &http.Server{
		Addr:              addr,
		Handler:           withLogging(mux),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Start serving
	log.Printf("listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}
