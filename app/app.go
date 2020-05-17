// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package app

import (
	"CiFa/service"
	"fmt"
	"sync"
)

type App struct {
	Conf    appConf
	Runtime appRuntime
}

/* 单例处理 */

var appInstance *App
var once sync.Once

func GetInstance() *App {
	//logging.Log(logging.INFO, logging.CallFromOuter+1, "Get app instance.")
	once.Do(func() {
		//logging.Info("App instance not created, new one...")
		appInstance = &App{}
		//logging.Info("App instance created at", fmt.Sprintf("%p", appInstance))
	})
	return appInstance
}

/* Conf */

type appConf struct {
	StaticDir     string `json:"static_dir"`      // 静态服务的文件目录
	TempDirPrefix string `json:"temp_dir_prefix"` // 临时文件目录的前缀
}

/* Runtime */

type appRuntime struct {
	Service *service.Service
}

func (a *App) Test() error {
	if a.Conf.TempDirPrefix == "" {
		return fmt.Errorf("TempDirPrefix Config Missing")
	}
	if a.Conf.StaticDir == "" {
		return fmt.Errorf("StaticDir Config Missing")
	}
	return nil
}

func (a *App) Run() {
	a.Runtime.Service = service.NewService(a.Conf.StaticDir, a.Conf.TempDirPrefix)
}
