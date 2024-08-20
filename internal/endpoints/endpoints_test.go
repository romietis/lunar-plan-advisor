package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/romietis/lunar-plan-advisor/v2/advisor"
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	return router
}

func TestEndpoint(t *testing.T) {
	router := SetUpRouter()
	router.LoadHTMLGlob("../../assets/templates/*")
	router.GET("/", Home)

	w := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, request)
	if w.Code != http.StatusOK {
		t.Errorf("wanted response code %v, got %v", http.StatusOK, w.Code)
	}
}

func TestPlansEndpointValidInput(t *testing.T) {
	router := SetUpRouter()
	router.GET("/plans", GetPlans)

	w := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/plans?balance=1000", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, request)

	if w.Code != http.StatusOK {
		t.Errorf("wanted response code %v, got %v", http.StatusOK, w.Code)
	}
	var expectedJsonStruct advisor.Best
	if err = json.Unmarshal(w.Body.Bytes(), &expectedJsonStruct); err != nil {
		t.Fatal(err)
	}
	if expectedJsonStruct.Plans[0].Name == "" {
		t.Error("expected non-empty string for a plan name")
	}
}

func TestPlansEndpointMissingBalance(t *testing.T) {
	router := SetUpRouter()
	router.GET("/plans", GetPlans)

	w := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/plans?balance", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, request)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wanted response code %v, got %v", http.StatusBadRequest, w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "balance is required") {
		t.Errorf("response body does not contain the expected error message")
	}
}

func TestPlansEndpointInvalidBalance(t *testing.T) {
	router := SetUpRouter()
	router.GET("/plans", GetPlans)

	w := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/plans?balance=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(w.Body)
	router.ServeHTTP(w, request)
	if w.Code != http.StatusBadRequest {
		t.Errorf("wanted response code %v, got %v", http.StatusBadRequest, w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "invalid input") {
		t.Errorf("response body does not contain the expected error message")
	}
}

func TestPlansEndpointNegativeInput(t *testing.T) {
	router := SetUpRouter()
	router.GET("/plans", GetPlans)

	w := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/plans?balance=-1000", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, request)
	if w.Code != http.StatusBadRequest {
		t.Errorf("wanted response code %v, got %v", http.StatusBadRequest, w.Code)
	}
}
