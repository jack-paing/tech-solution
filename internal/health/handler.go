package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Check if the application is up
func Check(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
