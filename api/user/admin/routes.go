package admin

import (
	"clientApi/config/dependency"

	"github.com/gin-gonic/gin"
)

func Router(rg *gin.RouterGroup) {
	business := NewAdminBusiness(dependency.CognitoAdmin)
	controller := NewAdminController(business)

	adminGroup := rg.Group("/admin")
	{
		adminGroup.POST("/create", controller.CreateUser)
	}

}
