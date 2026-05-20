package main

import (
	"net/http"

	"github.com/romietis/lunar-plan-advisor/v3/advisor"
	"github.com/romietis/lunar-plan-advisor/v3/internal/endpoints"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	// Static assets
	router.LoadHTMLGlob("assets/templates/*")
	router.Static("/css", "assets/css")
	router.Static("/js", "assets/js")

	// Google Search Console
	router.GET("/google0c4ea5396b01145c.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "google0c4ea5396b01145c.html", nil)
	})

	// Serve script
	router.GET("/", endpoints.Home)

	// Return the built-in default plan configuration
	router.GET("/plans", func(c *gin.Context) {
		endpoints.GetPlans(c, defaults)
	})

	// Calculate best plan(s) for a balance against the supplied (or default) plans
	router.POST("/plans/best", func(c *gin.Context) {
		endpoints.PostBestPlans(c, defaults)
	})

	router.Run()

}
