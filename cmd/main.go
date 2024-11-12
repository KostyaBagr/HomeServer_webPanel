// Entry point of the project

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/KostyaBagr/HomeServer_webPanel/pkg/settings"
	"github.com/KostyaBagr/HomeServer_webPanel/internal/routers"
)

func init() {
	settings.Setup()

}

func main() {
	gin.SetMode(settings.ServerSetting.RunMode)

	routersInit := routers.InitializeRoutes()
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", settings.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()


}