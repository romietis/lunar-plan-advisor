package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/romietis/lunar-plan-advisor/internal/endpoints"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(*godog.Scenario) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendRequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return
	}

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	c, router := gin.CreateTestContext(a.resp)
	router.LoadHTMLGlob("../../assets/templates/*")
	c.Request = req

	switch endpoint {
	case "/":
		endpoints.Home(c)
	default:
		err = fmt.Errorf("unknown endpoint: %s", endpoint)
	}
	return
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../requirements/"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse(sc)
		return ctx, nil
	})
	ctx.When(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, api.iSendRequestTo)
	ctx.Then(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
}
