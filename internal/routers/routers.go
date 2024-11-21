// this file contains server routes
package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/KostyaBagr/HomeServer_webPanel/initializers"
	"github.com/KostyaBagr/HomeServer_webPanel/internal/handlers"
	"github.com/KostyaBagr/HomeServer_webPanel/internal/middlewares"
)


func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

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
			middlewares.CheckAuth,
					handlers.ServerConfigurationHandler())
			systemGroup.POST("/manage/reboot", 
					middlewares.CheckAuth,
					handlers.RebootServerHandler())
			systemGroup.POST("/manage/poweroff", 
					middlewares.CheckAuth,
					handlers.PowerOffServerHandler())
		}
		componentsGroup := api.Group("components")
		{
			componentsGroup.GET("/cpu-info", 
				middlewares.CheckAuth, 
				handlers.GetCpuDetailInfo())
			componentsGroup.GET("/ram-info", 
				middlewares.CheckAuth, 
				handlers.GetRamDetailInfo())
			componentsGroup.GET("/disks-info", 
				// middlewares.CheckAuth, 
				handlers.GetDiskDetailInfo())
		}
		
		userGroup := api.Group("user")
		{
			userGroup.POST("/auth/register", handlers.CreateUser())
			userGroup.POST("auth/login", handlers.Login())
			userGroup.GET("/profile", middlewares.CheckAuth, 
			handlers.GetUserProfile())

		
		}
	}
	return r 
}