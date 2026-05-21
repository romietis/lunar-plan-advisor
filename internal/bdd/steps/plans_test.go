package steps

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/romietis/lunar-plan-advisor/v3/advisor"
	"github.com/romietis/lunar-plan-advisor/v3/internal/endpoints"
)

type apiContext struct {
	balance  float64
	response *httptest.ResponseRecorder
}

func (ac *apiContext) resetResponse(*godog.Scenario) {
	ac.response = httptest.NewRecorder()
}

func (ac *apiContext) givenBalance(balance string) error {
	floatValue, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		return err
	}
	ac.balance = floatValue
	return nil
}

func (a *apiContext) sendRequestTo(method, endpoint string) error {
	defaults := advisor.PlansConfig{
		Plans: []advisor.PlanConfig{
			{Name: "Light", AnnualInterestRate: 1.25, Fee: 0.0, Cap: 100000},
			{Name: "Standard", AnnualInterestRate: 1.5, Fee: 29.0, Cap: 100000},
			{Name: "Plus", AnnualInterestRate: 1.75, Fee: 69.0, Cap: 0},
			{Name: "Unlimited", AnnualInterestRate: 2.25, Fee: 139.0, Cap: 0},
		},
	}

	body := fmt.Sprintf(`{"balance":%v}`, a.balance)
	req, err := http.NewRequest(method, endpoint, strings.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	h := &endpoints.Handler{Defaults: defaults}
	mux := http.NewServeMux()
	mux.HandleFunc(method+" "+endpoint, h.PostBestPlans)

	mux.ServeHTTP(a.response, req)
	return nil

}

func (ac *apiContext) responseCodeShouldBe(code int) error {
	if ac.response.Code != code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, ac.response.Code)
	}
	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(s *godog.ScenarioContext) {
	api := &apiContext{}

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse(sc)
		return ctx, nil
	})

	s.Step(`^a blance of (-?\d+(\.\d+)?([eE][-+]?\d+)?)\s*DKK$`, api.givenBalance)
	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, api.sendRequestTo)
	s.Step(`^the response code should be (\d+)$`, api.responseCodeShouldBe)
}
