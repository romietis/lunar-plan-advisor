package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romietis/lunar-plan-advisor/advisor"
)

func main() {
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
	router.GET("/", advisor.Endpoint)

	// Public endpoint to fetch plans
	router.GET("/plans", advisor.PlansEndpoint)

	router.Run()

}
