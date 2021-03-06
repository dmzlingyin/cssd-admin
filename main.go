package main

import (
	"fmt"
	"net/http"

	"cssd-admin/models"
	"cssd-admin/pkg/setting"
	"cssd-admin/routers"
)

func main() {
	router := routers.InitRouter()
	models.Init()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
