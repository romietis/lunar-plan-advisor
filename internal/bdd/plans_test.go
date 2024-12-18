package bdd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
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
	planConfig := advisor.Plans{
		Plans: []advisor.Plan{
			{Name: "Light", AnnualInterestRate: 1.25, Fee: 0.0, Cap: 100000},
			{Name: "Standard", AnnualInterestRate: 1.5, Fee: 29.0, Cap: 100000},
			{Name: "Plus", AnnualInterestRate: 1.75, Fee: 69.0, Cap: 0},
			{Name: "Unlimited", AnnualInterestRate: 2.25, Fee: 139.0, Cap: 0},
		},
	}

	appendedQuearyParam := fmt.Sprintf("%v?balance=%v", endpoint, a.balance)
	req, err := http.NewRequest(method, appendedQuearyParam, nil)
	if err != nil {
		return err
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET(endpoint, func(c *gin.Context) {
		endpoints.GetPlans(c, planConfig)
	})

	router.ServeHTTP(a.response, req)
	return nil

}

func (ac *apiContext) responseCodeShouldBe(code int) error {
	if ac.response.Code != code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, ac.response.Code)
	}
	return nil
}

func (ac *apiContext) responseShouldMatch() error {
	var expectedJsonStruct advisor.Plans
	if err := json.Unmarshal(ac.response.Body.Bytes(), &expectedJsonStruct); err != nil {
		return err
	}

	for _, plan := range expectedJsonStruct.Plans {
		if plan.Name == "" {
			return errors.New("missing required value")
		}
	}

	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
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
	s.Step(`^the response should match json`, api.responseShouldMatch)
}
