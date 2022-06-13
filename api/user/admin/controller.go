package admin

import (
	"github.com/gin-gonic/gin"
)

type IAdminController interface {
	CreateUser(c *gin.Context)
}

type AdminController struct {
	adminBusiness IAdminBusiness
}

func NewAdminController(adminBusiness IAdminBusiness) IAdminController {
	return &AdminController{adminBusiness}
}

func (a *AdminController) CreateUser(ctx *gin.Context) {
	statusCode, err := a.adminBusiness.CreateUser()
	if err != nil {
		ctx.AbortWithStatusJSON(statusCode, map[string]string{"code": "error", "message": err.Error()})
	}

	ctx.AbortWithStatusJSON(statusCode, map[string]string{"code": "success", "message": "success"})
}
