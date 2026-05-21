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

	h := &endpoints.Handler{
		Defaults:  defaults,
		Templates: tmpl,
	}

	mux := http.NewServeMux()

	mux.Handle("GET /css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css"))))
	mux.Handle("GET /js/", http.StripPrefix("/js/", http.FileServer(http.Dir("assets/js"))))

	mux.HandleFunc("GET /{$}", h.Home)
	mux.HandleFunc("GET /plans", h.GetPlans)
	mux.HandleFunc("POST /plans/best", h.PostBestPlans)
	mux.HandleFunc("GET /google0c4ea5396b01145c.html", serveTemplate(tmpl, "google0c4ea5396b01145c.html"))

	addr := ":" + port()
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(tmpl *template.Template, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, name, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}
