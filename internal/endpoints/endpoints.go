package endpoints

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/romietis/lunar-plan-advisor/v3/advisor"
)

// Handler holds the dependencies shared by the HTTP endpoints.
type Handler struct {
	Defaults  advisor.PlansConfig
	Templates *template.Template
}

type bestPlansRequest struct {
	Balance *float64             `json:"balance"`
	Plans   []advisor.PlanConfig `json:"plans"`
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.Templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetPlans returns the built-in default plan configuration. The UI uses this
// to seed first-time visitors and as the "reset to defaults" source.
func (h *Handler) GetPlans(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.Defaults)
}

// PostBestPlans calculates the best plan(s) for a given balance against a
// plan configuration supplied in the request body. When Plans is omitted, the
// server-side defaults are used so the endpoint stays usable without a config.
func (h *Handler) PostBestPlans(w http.ResponseWriter, r *http.Request) {
	var req bestPlansRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if req.Balance == nil {
		writeError(w, http.StatusBadRequest, "balance is required")
		return
	}

	config := advisor.PlansConfig{Plans: req.Plans}
	if len(config.Plans) == 0 {
		config = h.Defaults
	} else if err := config.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	bestPlans, err := config.CalculatePlans(*req.Balance)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, advisor.Plans{Plans: bestPlans})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
