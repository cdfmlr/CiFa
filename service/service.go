// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"CiFa/util/logging"
	"net/http"
)

type Service struct {
	WordFaSessionHolder *WordFaSessionHolder
	StaticDir           string
	TempDirPrefix       string

	fileServer http.Handler
}

func NewService(staticDir string, tempDirPrefix string) *Service {
	s := &Service{
		WordFaSessionHolder: NewWordFaSessionHolder(),
		StaticDir:           staticDir,
		TempDirPrefix:       tempDirPrefix,
	}
	//s.fileServer = http.StripPrefix("/static", http.FileServer(http.Dir(s.StaticDir)))
	s.fileServer = http.FileServer(http.Dir(s.StaticDir))
	return s
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logging.Info("HTTP Serve: ", r.Method, r.URL.Path)

	switch r.URL.Path {
	case "/api/wordfa":
		s.ApiWordfa(w, r)
	case "/api/sort/float":
		s.ApiSortFloat(w, r)
	case "/api/strsearch":
		s.ApiStrsearch(w, r)
	default:
		// 对于其他 URL Path，使用 StaticDir 上的文件服务
		// 例如: GET /index.html 返回文件 $StaticDir/index.html
		s.fileServer.ServeHTTP(w, r)
	}
}
