package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/romietis/lunar-plan-advisor/v3/advisor"
	"github.com/romietis/lunar-plan-advisor/v3/internal/endpoints"
)

func main() {
	defaults := advisor.PlansConfig{
		Plans: []advisor.PlanConfig{
			{Name: "Light", AnnualInterestRate: 0.75, Fee: 0.0, Cap: 100000},
			{Name: "Standard", AnnualInterestRate: 1.0, Fee: 29.0, Cap: 100000},
			{Name: "Plus", AnnualInterestRate: 1.25, Fee: 69.0, Cap: 0},
			{Name: "Unlimited", AnnualInterestRate: 1.75, Fee: 139.0, Cap: 0},
		},
	}

	tmpl := template.Must(template.ParseGlob("assets/templates/*"))

	handlers := &endpoints.Handler{
		Defaults:  defaults,
		Templates: tmpl,
	}

	mux := http.NewServeMux()

	mux.Handle("GET /css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css"))))
	mux.Handle("GET /js/", http.StripPrefix("/js/", http.FileServer(http.Dir("assets/js"))))

	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /plans", handlers.GetPlans)
	mux.HandleFunc("POST /plans/best", handlers.PostBestPlans)

	addr := ":" + port()
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}
