// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"CiFa/util/logging"
	"CiFa/util/strsearch"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ApiStrsearch 处理 POST /api/strsearch, 在给定字符串做子串搜索
// Request:
//		POST /api/strsearch
// 		Body: JSON:
//			{"algorithm": 0, "text": "abcbab", "pattern": "ab"}
//				text: 	 : string: 父字符串，在此字符串中搜索子串 pattern
//				pattern	 : string: 子字符串，在 text 中搜索此字符串
//				algorithm: int:    字符串搜索算法，0~3, 分别是:
//											regexp.FindAllIndex (go lib)，KMP 算法，Rabin-Karp 算法，暴力法
// Response:
//		Success: JSON: {"index": [0, 4], "time_cost": "time cost"}	// index 是 pattern 在 text 中出现位置的索引，注意中文字符不是"第几个字"！
//		Error:   JSON: {"error": "error description"}
func (s *Service) ApiStrsearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseJson(&w, ErrorResponse{ErrorDescription: "Request should be POST"})
		return
	}
	defer r.Body.Close()

	var body apiStrsearchRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logging.Error(
			fmt.Sprintf("ApiStrsearch failed: Cannot Decode Body Json:\n\t|-err: %v\n\t|-req: %#v",
				err, r,
			),
		)
		responseJson(&w, ErrorResponse{ErrorDescription: err.Error()})
		return
	}
	if body.Algorithm < 0 || body.Algorithm > 3 {
		body.Algorithm = strsearch.LibRe
	}
	start := time.Now()
	index := strsearch.By(body.Algorithm).FindAll(body.Text, body.Pattern)
	elapsed := time.Since(start)
	logging.Info(fmt.Sprintf("ApiSortFloat success: %#v", body))
	responseJson(&w, PostApiStrsearchResponse{
		Index:    index,
		TimeCost: fmt.Sprintf("%v", elapsed),
	})
}

type apiStrsearchRequestBody struct {
	Text      string `json:"text"`
	Pattern   string `json:"pattern"`
	Algorithm int    `json:"algorithm"`
}
