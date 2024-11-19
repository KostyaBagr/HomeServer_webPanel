package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/KostyaBagr/HomeServer_webPanel/internal/services"
)

func ServerConfigurationHandler() gin.HandlerFunc {
	// Return server configuration info
	return func(c *gin.Context) {
		system, err := services.ServerConfiguration()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not retrieve server config information",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": system,
		})
	}
}


func RebootServerHandler() gin.HandlerFunc{
	// Call fucntion for reboot server
	return func (c *gin.Context)  {
		reboot, err := services.RebootServer()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not reboot the server",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": reboot,
		})
	}
}

func PowerOffServerHandler() gin.HandlerFunc{
	// Call function for shutdown server
	return func(c *gin.Context) {
		poweroff, err := services.PowerOffserver()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not poweroff the server",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": poweroff,
		})
	}
}