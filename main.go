package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/romietis/lunar-plan-advisor/advisor"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("assets/templates/*")
	router.Static("/css", "assets/css")
	router.Static("/js", "assets/js")

	router.GET("/", advisor.Endpoint)
	router.GET("/google0c4ea5396b01145c.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "google0c4ea5396b01145c.html", nil)
	})

	router.Run()

}
