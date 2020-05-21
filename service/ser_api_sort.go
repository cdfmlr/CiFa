// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"CiFa/util/logging"
	"CiFa/util/sortalgo"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// apiWordfaPost 处理 POST /api/sort/float, 对给定浮点数序列进行排序
// Request:
//		POST /api/sort/float
// 		Body: JSON:
//			{"algorithm": 0, "data": [2, 1, 3.0, 7, 4.4, ...]}
//				data	 : []float: 要排序的数据
//				algorithm: int    : 指定排序算法，0~8, 分别是:
//									sort.Sort (go lib)，sort.Stable (go lib)，快速排序，堆排序，
//									归并排序，希尔排序，希尔排序(并发), 插入排序，选择排序
// Response:
//		Success: JSON: {"result": [1, 2, 3.0, 4.4, 7, ...], "time_cost": "time cost"}
//		Error:   JSON: {"error": "error description"}
func (s *Service) ApiSortFloat(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseJson(&w, ErrorResponse{ErrorDescription: "Request should be POST"})
		return
	}
	defer r.Body.Close()

	var body apiSortFloatRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logging.Error(
			fmt.Sprintf("ApiSortFloat failed: Cannot Decode Body Json:\n\t|-err: %v\n\t|-req: %#v",
				err, r,
			),
		)
		responseJson(&w, ErrorResponse{ErrorDescription: err.Error()})
		return
	}

	start := time.Now()
	if len(body.Data) > 1 {
		if body.Algorithm < 0 || body.Algorithm > 8 {
			body.Algorithm = sortalgo.StlSort
		}
		sortalgo.By(body.Algorithm).Sort(body.Data)
	}
	elapsed := time.Since(start)
	logging.Info(fmt.Sprintf("ApiSortFloat success: %#v", body))
	responseJson(&w, PostApiSortFloatResponse{
		Result:   body.Data,
		TimeCost: fmt.Sprintf("%v", elapsed),
	})
}

type apiSortFloatRequestBody struct {
	Algorithm int      `json:"algorithm"`
	Data      float64S `json:"data"`
}

type float64S []float64

func (f float64S) Len() int {
	return len(f)
}

func (f float64S) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f float64S) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
