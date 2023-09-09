package main

import (
	"fmt"
	"github.com/422511186/gin_Blog_demo/models"
	"github.com/422511186/gin_Blog_demo/pkg/gredis"
	"github.com/422511186/gin_Blog_demo/pkg/logging"
	"github.com/422511186/gin_Blog_demo/pkg/setting"
	"github.com/422511186/gin_Blog_demo/pkg/util"
	"github.com/422511186/gin_Blog_demo/routers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
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
