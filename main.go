package main

import (
	"github.com/gin-gonic/gin"
	"github.com/romietis/lunar-plan-advisor/advisor"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("assets/templates/*")
	router.Static("/css", "assets/css")
	router.Static("/js", "assets/js")

	router.GET("/", advisor.Endpoint)

	router.Run()

}
