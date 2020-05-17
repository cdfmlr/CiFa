// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"CiFa/wordfa"
	"encoding/json"
	"net/http"
)

// 错误时的返回模版
type ErrorResponse struct {
	ErrorDescription string `json:"error"`
}

// POST api/wordfa 成功的返回
type PostApiWordfaResponse struct {
	Success string `json:"success"`
}

// GET api/wordfa 成功的返回
type GetApiWordfaResponse struct {
	Progress float32       `json:"progress"`
	Result   wordfa.Result `json:"result"`
}

// responseJson 将传过来的 resp Marshal 成 Json，写到 w
func responseJson(w *http.ResponseWriter, resp interface{}) {
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	if _, err = (*w).Write(js); err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
	}
}
