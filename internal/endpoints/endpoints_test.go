package endpoints

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	return router
}

func TestEndpoint(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a request to pass to our handler
	request, _ := http.NewRequest("GET", "/", nil)

	// Create a Gin context and engine (router := gin.Default()) from the response recorder
	c, router := gin.CreateTestContext(w)

	// Load HTML templates
	router.LoadHTMLGlob("../../assets/templates/*")

	c.Request = request

	// Call the endpoint, check the status code
	Home(c)
	if w.Code != http.StatusOK {
		t.Errorf("response code is %v", w.Code)
	}
}

func TestPlansEndpointValidInput(t *testing.T) {
	r := SetUpRouter()
	r.GET("/plans", GetPlans)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/plans?balance=1000", nil)

	r.ServeHTTP(w, request)

	if w.Code != http.StatusOK {
		t.Errorf("response code is %v", w.Code)
	}
}

func TestPlansEndpointMissingBalance(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	GetPlans(c)
	if w.Code != http.StatusBadRequest {
		t.Errorf("response code is %v", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "balance is required") {
		t.Errorf("response body does not contain the expected error message")
	}
}

func TestPlansEndpointInvalidBalance(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/?balance=invalid", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	GetPlans(c)
	if w.Code != http.StatusBadRequest {
		t.Errorf("response code is %v", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "invalid input") {
		t.Errorf("response body does not contain the expected error message")
	}
}

func TestPlansEndpointNegativeInput(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/?balance=-1000", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	GetPlans(c)
	if w.Code != http.StatusBadRequest {
		t.Errorf("response code is %v", w.Code)
	}
}
