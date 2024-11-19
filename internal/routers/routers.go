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
		systemGroup := api.Group("server") 
		{
			systemGroup.GET("/system-info", 
					handlers.ServerConfigurationHandler())
			systemGroup.GET("/manage/reboot", 
					handlers.RebootServerHandler())
			systemGroup.GET("/manage/poweroff", 
					handlers.PowerOffServerHandler())
		}
		componentsGroup := api.Group("components")
		{
			componentsGroup.GET("/cpu-info", handlers.GetCpuDetailInfo())
			componentsGroup.GET("/ram-info", handlers.GetRamDetailInfo())
			componentsGroup.GET("/disks-info", handlers.GetDiskDetailInfo())
		}
	}
	return r 
}