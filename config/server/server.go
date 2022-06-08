package server

import (
	"clientApi/api"
	docs "clientApi/docs"

	"clientApi/config/env"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	gin.SetMode(env.GinMode)
	r := gin.New()

	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/"))

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
