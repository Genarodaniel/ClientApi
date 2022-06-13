package api

import (
	"clientApi/api/healthCheck"
	"clientApi/api/user/admin"

	"github.com/gin-gonic/gin"
)

// @Version 1.0
func Router(e *gin.Engine) {
	healthCheck.Router(e.Group("/"))

	userGroup := e.Group("/user")
	{
		admin.Router(userGroup)
	}
}
