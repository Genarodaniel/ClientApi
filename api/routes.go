package api

import (
	"clientApi/api/healthCheck"

	"github.com/gin-gonic/gin"
)

// @Version 1.0
func Router(e *gin.Engine) {
	healthCheck.Router(e.Group("/"))
}
