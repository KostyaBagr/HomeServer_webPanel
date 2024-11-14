// This file contains handlers for cpu, ram, disk and so on
package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/KostyaBagr/HomeServer_webPanel/internal/services"
)

func GetCpuDetailInfo() gin.HandlerFunc {
	// Return CPU load info
	return func(c *gin.Context) {
		cpu, err := services.CpuDetailInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not retrieve CPU information",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": cpu,
		})
	}
}