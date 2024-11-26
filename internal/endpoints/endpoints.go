package endpoints

import (
	"net/http"
	"strconv"

	"github.com/romietis/lunar-plan-advisor/v3/advisor"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetPlans(c *gin.Context, planConfig advisor.Plans) {

	plans := advisor.Plans{Plans: planConfig.Plans}

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

	bestPlans, err := plans.CalculatePlans(balance_float)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	best := advisor.Plans{
		Plans: bestPlans,
	}
	c.JSON(http.StatusOK, best)

}
