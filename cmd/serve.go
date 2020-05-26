/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"CiFa/app"
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var port int
var tempDirPrefix string
var staticDir string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a CiFa web serve",
	Long:  `Start a CiFa web serve and then you can use CiFa in your browser for <words frequency analyzing>, <sort algorithm test> and <string matching test>.`,
	Run: func(cmd *cobra.Command, args []string) {
		cifa := app.GetInstance()

		if staticDir != "" {
			cifa.Conf.StaticDir = staticDir
		}

		if tempDirPrefix != "" {
			cifa.Conf.TempDirPrefix = tempDirPrefix
		}

		// 检查 app 配置完备性
		if err := cifa.Test(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Config Error:", err)
			flag.Usage()
			os.Exit(-1)
		}

		// 启动 app，监听服务
		_, _ = fmt.Fprintf(os.Stderr, "CiFa running on localhost:%v.\nstatic serve file on: %v\n", port, cifa.Conf.StaticDir)
		cifa.Run()
		err := http.ListenAndServe(fmt.Sprintf(":%v", port), cifa.Runtime.Service)

		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "http.ListenAndServe error:", err)
			os.Exit(-1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntVarP(&port, "port", "p", 9001, "`port` for service")
	serveCmd.Flags().StringVarP(&tempDirPrefix, "temp_dir_prefix", "t", "temp.cifa.", "name `prefix` for temp files' dir")
	serveCmd.Flags().StringVarP(&staticDir, "static_dir", "s", "./static", "static (web ui) `dist` path")
}
