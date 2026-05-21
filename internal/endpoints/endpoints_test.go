package endpoints

import (
	"encoding/json"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/romietis/lunar-plan-advisor/v3/advisor"
)

func defaultPlans() advisor.PlansConfig {
	return advisor.PlansConfig{
		Plans: []advisor.PlanConfig{
			{Name: "Light", AnnualInterestRate: 1.25, Fee: 0.0, Cap: 100000},
			{Name: "Standard", AnnualInterestRate: 1.5, Fee: 29.0, Cap: 100000},
			{Name: "Plus", AnnualInterestRate: 1.75, Fee: 69.0, Cap: 0},
			{Name: "Unlimited", AnnualInterestRate: 2.25, Fee: 139.0, Cap: 0},
		},
	}
}

func newTestHandler() *Handler {
	tmpl := template.Must(template.ParseGlob("../../assets/templates/*"))
	return &Handler{
		Defaults:  defaultPlans(),
		Templates: tmpl,
	}
}

func TestHome(t *testing.T) {
	h := newTestHandler()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h.Home(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("wanted response code %v, got %v", http.StatusOK, w.Code)
	}
}

func TestGetPlansReturnsDefaults(t *testing.T) {
	h := newTestHandler()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/plans", nil)
	h.GetPlans(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("wanted response code %v, got %v", http.StatusOK, w.Code)
	}
	var got advisor.Plans
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Plans) != len(h.Defaults.Plans) {
		t.Errorf("wanted %d plans, got %d", len(h.Defaults.Plans), len(got.Plans))
	}
	if got.Plans[0].Name != "Light" {
		t.Errorf("wanted first plan Light, got %s", got.Plans[0].Name)
	}
}

func postBest(t *testing.T, body string) *httptest.ResponseRecorder {
	t.Helper()
	h := newTestHandler()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/plans/best", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	h.PostBestPlans(w, req)
	return w
}

func TestPostBestPlansWithDefaults(t *testing.T) {
	w := postBest(t, `{"balance":1000}`)
	if w.Code != http.StatusOK {
		t.Errorf("wanted %v, got %v body=%s", http.StatusOK, w.Code, w.Body.String())
	}
	var got advisor.Plans
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Plans) == 0 || got.Plans[0].Name == "" {
		t.Errorf("expected at least one named best plan, got %+v", got)
	}
}

func TestPostBestPlansWithCustomPlans(t *testing.T) {
	body := `{"balance":1000,"plans":[{"name":"Custom","annualInterestRate":3.0,"fee":0,"cap":0}]}`
	w := postBest(t, body)
	if w.Code != http.StatusOK {
		t.Errorf("wanted %v, got %v body=%s", http.StatusOK, w.Code, w.Body.String())
	}
	var got advisor.Plans
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Plans) != 1 || got.Plans[0].Name != "Custom" {
		t.Errorf("expected Custom plan, got %+v", got.Plans)
	}
}

func TestPostBestPlansMissingBalance(t *testing.T) {
	w := postBest(t, `{}`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("wanted %v, got %v", http.StatusBadRequest, w.Code)
	}
	if !strings.Contains(w.Body.String(), "balance is required") {
		t.Errorf("unexpected body: %s", w.Body.String())
	}
}

func TestPostBestPlansNegativeBalance(t *testing.T) {
	w := postBest(t, `{"balance":-1000}`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("wanted %v, got %v", http.StatusBadRequest, w.Code)
	}
}

func TestPostBestPlansInvalidJSON(t *testing.T) {
	w := postBest(t, `not json`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("wanted %v, got %v", http.StatusBadRequest, w.Code)
	}
}

func TestPostBestPlansInvalidPlans(t *testing.T) {
	body := `{"balance":1000,"plans":[{"name":"","annualInterestRate":1.0}]}`
	w := postBest(t, body)
	if w.Code != http.StatusBadRequest {
		t.Errorf("wanted %v, got %v", http.StatusBadRequest, w.Code)
	}
	if !strings.Contains(w.Body.String(), "name can't be empty") {
		t.Errorf("unexpected body: %s", w.Body.String())
	}
}
