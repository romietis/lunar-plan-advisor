package main

import (
	"net/http"

	"github.com/romietis/lunar-plan-advisor/v2/advisor"
	"github.com/romietis/lunar-plan-advisor/v2/internal/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	planConfig := advisor.Plans{
		Plans: []advisor.Plan{
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

	// Public endpoint to fetch plans
	router.GET("/plans", func(c *gin.Context) {
		endpoints.GetPlans(c, planConfig)
	})

	router.Run()

}
