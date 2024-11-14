// this file contains server routes
package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/KostyaBagr/HomeServer_webPanel/internal/handlers"

)


func InitializeRoutes() *gin.Engine {
	// Function for initializing routes.
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("api/v1")
	{
		systemGroup := api.Group("system") 
		{
			systemGroup.GET("/system-info", handlers.ServerConfigurationHandler())
		}
	}
	return r 
}