package endpoints

import (
	"net/http"
	"strconv"

	"github.com/romietis/lunar-plan-advisor/advisor"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetPlans(c *gin.Context) {
	planConfig := []advisor.Plan{
		{Name: "Light", AnnualInterestRate: 1.25, Fee: 0.0, Cap: 100000},
		{Name: "Standard", AnnualInterestRate: 1.5, Fee: 29.0, Cap: 100000},
		{Name: "Plus", AnnualInterestRate: 1.75, Fee: 69.0, Cap: 0},
		{Name: "Unlimited", AnnualInterestRate: 2.25, Fee: 119.0, Cap: 0},
	}

	balance := c.Query("balance")
	if balance == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "balance is required",
		})
		return
	}

	balance_float, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	bestPlans, err := advisor.CalculatePlans(balance_float, planConfig)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"plans": bestPlans,
	})

}
