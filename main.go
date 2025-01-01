package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-programming-tour-book/blog_service/global"
	"github.com/go-programming-tour-book/blog_service/internal/routers"
	"github.com/go-programming-tour-book/blog_service/pkg/setting"
)

// 程序执行顺序： 全局变量初始化--> init方法 --> main方法
//  不要滥用 init方法， 如果init方法过多， 则很容易迷失在各个库的init方法

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
