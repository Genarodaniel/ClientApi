package server

import (
	"clientApi/api"
	docs "clientApi/docs"
	"fmt"

	"clientApi/config/dependency"
	"clientApi/config/env"
	"clientApi/config/recovery"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	env.Load()

	if err := dependency.Load(); err != nil {
		fmt.Println(err)
	}

	gin.SetMode(env.GinMode)
	r := gin.New()

	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/"))
	r.Use(gin.CustomRecovery(recovery.Filter))
	api.Router(r)

	// Configure Swagger
	ConfigSwagger(r)

	return r
}

func ConfigSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "Swagger Client API"
	docs.SwaggerInfo.Description = "API responsible for managing Clients Repository"
	docs.SwaggerInfo.Version = "1.0"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
