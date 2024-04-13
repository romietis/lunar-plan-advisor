package advisor

import (
	"net/http"
	"net/http/httptest"
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
