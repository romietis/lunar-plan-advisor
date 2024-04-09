package advisor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Endpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Plan",
	})
}
