// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"net/http"
	"strings"
)

type Service struct {
	WordFaSessionHolder *WordFaSessionHolder
	TempDirPrefix       string
}

func NewService(tempDirPrefix string) *Service {
	return &Service{
		WordFaSessionHolder: NewWordFaSessionHolder(),
		TempDirPrefix:       tempDirPrefix,
	}
}

func (s *Service) RegisterHandles(baseUrl string) {
	if strings.TrimSpace(baseUrl) == "" {
		baseUrl = "/api"
	}
	http.HandleFunc(baseUrl+"/wordfa", s.ApiWordfa)
}
