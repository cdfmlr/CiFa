// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"CiFa/app"
	"CiFa/util"
	"CiFa/util/logging"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 命令行参数
	flag.Usage = usage

	confFile := flag.String("c", "", "set the configuration `file`")
	port := flag.Uint("p", 9001, "set the `port` for service")
	staticDir := flag.String("s", "./static", "set the `static_dir_path`")
	tempDirPrefix := flag.String("t", "temp.cifa.", "set the `temp_dir_prefix`")

	flag.Parse()

	// 配置 app，配置文件优先级高于命令行参数
	cifa := app.GetInstance()

	if *staticDir != "" {
		cifa.Conf.StaticDir = *staticDir
	}

	if *tempDirPrefix != "" {
		cifa.Conf.TempDirPrefix = *tempDirPrefix
	}

	if *confFile != "" {
		if err := util.LoadJsonFile(*confFile, &cifa.Conf); err != nil {
			logging.Error("Failed to load configuration file:", err)
		}
	}

	// 检查 app 配置完备性
	if err := cifa.Test(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Config Error:", err)
		flag.Usage()
		os.Exit(-1)
	}

	// 启动 app，监听服务
	_, _ = fmt.Fprintf(os.Stderr, "CiFa running on localhost:%v.\n static file serve: %v\n", *port, *staticDir)
	cifa.Run()
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), cifa.Runtime.Service)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "http.ListenAndServe error:", err)
		os.Exit(-1)
	}
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `
CiFa  v0.0.1
------------------------------
			Powered by CDFMLR

Usage: cifa [-s static_dir] [-p port] [-t temp_dir_prefix] [-c config_file]

Options:
`)
	flag.PrintDefaults()
}
