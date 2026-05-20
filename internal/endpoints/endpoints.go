package endpoints

import (
	"net/http"

	"github.com/romietis/lunar-plan-advisor/v3/advisor"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// GetPlans returns the built-in default plan configuration. The UI uses this
// to seed first-time visitors and as the "reset to defaults" source.
func GetPlans(c *gin.Context, defaults advisor.PlansConfig) {
	c.JSON(http.StatusOK, defaults)
}

type bestPlansRequest struct {
	Balance *float64             `json:"balance"`
	Plans   []advisor.PlanConfig `json:"plans"`
}

// PostBestPlans calculates the best plan(s) for a given balance against a
// plan configuration supplied in the request body. When Plans is omitted, the
// server-side defaults are used so the endpoint stays usable without a config.
func PostBestPlans(c *gin.Context, defaults advisor.PlansConfig) {
	var req bestPlansRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Balance == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "balance is required"})
		return
	}

	config := advisor.PlansConfig{Plans: req.Plans}
	if len(config.Plans) == 0 {
		config = defaults
	} else if err := config.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bestPlans, err := config.CalculatePlans(*req.Balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, advisor.Plans{Plans: bestPlans})
}
