package advisor

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEndpoint(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a request to pass to our handler
	request, _ := http.NewRequest("GET", "/", nil)

	// Create a Gin context and engine (router := gin.Default()) from the response recorder
	c, router := gin.CreateTestContext(w)

	// Load HTML templates
	router.LoadHTMLGlob("../assets/templates/*")

	c.Request = request

	// Call the endpoint, check the status code
	Endpoint(c)
	if w.Code != http.StatusOK {
		t.Errorf("response code is %v", w.Code)
	}
}

func TestPlansEndpoint_ValidInput(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/?balance=1000", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	PlansEndpoint(c)
	if w.Code != http.StatusOK {
		t.Errorf("response code is %v", w.Code)
	}
}

func TestPlansEndpoint_MissingBalance(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	PlansEndpoint(c)
	if w.Code != http.StatusBadRequest {
		t.Errorf("response code is %v", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "balance is required") {
		t.Errorf("response body does not contain the expected error message")
	}
}

func TestPlansEndpoint_InvalidBalance(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/?balance=invalid", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	PlansEndpoint(c)
	if w.Code != http.StatusBadRequest {
		t.Errorf("response code is %v", w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "invalid input") {
		t.Errorf("response body does not contain the expected error message")
	}
}

func TestPlansEndpoint_NegativeInput(t *testing.T) {
	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/?balance=-1000", nil)
	c, _ := gin.CreateTestContext(w)
	c.Request = request
	PlansEndpoint(c)
	if w.Code != http.StatusBadRequest {
		t.Errorf("response code is %v", w.Code)
	}
}
