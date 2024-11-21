// This file contains handlers for cpu, ram, disk and so on
package handlers

import (

	"net/http"

	"github.com/KostyaBagr/HomeServer_webPanel/internal/services"
	"github.com/gin-gonic/gin"
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

func GetRamDetailInfo() gin.HandlerFunc {
	// Return Ram info
	return func(c *gin.Context) {
		ram, err := services.RamDetailInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not retrieve Ram information",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": ram,
		})
	}
}


func GetDiskDetailInfo() gin.HandlerFunc {
	// Return Disk info
	return func(c *gin.Context) {
		disk, err := services.DiskUsageSummary()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not retrieve disk information",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result":  disk,
		})
	}
}

