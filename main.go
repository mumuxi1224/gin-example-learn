package main

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers"
	"log"
	"net/http"
)

func main() {
	//router:=gin.Default()
	//router.GET("/test", func(context *gin.Context) {
	//	context.JSON(200,gin.H{
	//		"message":"test",
	//	})
	//})
	router:=routers.InitRouter()

	s:=&http.Server{
		Addr: fmt.Sprintf(":%d",setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,//2的20次方 1024*1024
	}

	log.Fatal( s.ListenAndServe())
}