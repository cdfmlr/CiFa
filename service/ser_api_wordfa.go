// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package service

import (
	"CiFa/util"
	"CiFa/util/logging"
	"CiFa/util/sortalgo"
	"CiFa/util/strsearch"
	"CiFa/wordfa"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

// ApiWordfa 接收 /api/wordfa 的请求，并根据请求方式分发给特定函数进行处理
// GET  -> apiWordfaGet
// POST -> apiWordfaPost
func (s *Service) ApiWordfa(w http.ResponseWriter, r *http.Request) {
	// 解析 Form
	if err := r.ParseForm(); err != nil {
		logging.Error("ApiWordfa failed: ParseForm Error:", err)
		responseJson(&w, ErrorResponse{ErrorDescription: "Cannot Parse Form"})
		return
	}
	_ = r.ParseMultipartForm(32 << 20)
	// 验证 token
	if token := r.FormValue("token"); token == "" {
		logging.Warning("ApiWordfa failed: bad token")
		responseJson(&w, ErrorResponse{ErrorDescription: "Bad Token!"})
		return
	}
	// 分发处理
	switch r.Method {
	case "POST":
		s.apiWordfaPost(w, r)
	case "GET":
		s.apiWordfaGet(w, r)
	}
}

// apiWordfaPost 处理 GET /api/wordfa, 获取 wordfa 会话任务的处理进度/结果
// Request:
//		GET /api/wordfa
// 		Form:
//			token :FormValue string: 识别客户端身份的 token
// Response:
//		Task Running:  JSON: {"progress": 0.7}
//		Task Finished: JSON: {"progress": 1.0, "result": [{"keyword": 26}, {...}, ...]}
//		Error:         JSON: {"error": "error description"}
func (s *Service) apiWordfaGet(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	session, ok := s.WordFaSessionHolder.Get(token)
	if !ok {
		logging.Warning("apiWordfaGet failed: session not exist")
		responseJson(&w, ErrorResponse{ErrorDescription: "session not exist"})
		return
	}

	progress := session.Task.GetProgress()
	var result wordfa.Result
	if progress >= 1 {
		result, _ = session.Task.GetResult(session.SortAlgorithm)
	}

	logging.Info(fmt.Sprintf(
		"apiWordfaGet success: token=%#v\n\t--> progress: %v\n\t--> result: %v",
		token, progress, result,
	))

	responseJson(&w, GetApiWordfaResponse{
		Progress: progress,
		Result:   result,
	})

}

// apiWordfaPost 处理 POST /api/wordfa, 新建 wordfa 任务会话
// Request:
//		POST /api/wordfa
// 		Form:
//			token		:FormValue string: 识别客户端身份的 token
//			keywords	:FormValue string: 要检测的关键词，多个词间用逗号(',' 或 '，')隔开
//			file		:FormFile  file:   要检测的文件，单个文本文件(text/plain)，或多个文件的 zip 打包(application/zip)
//			sort_by		:FormValue int:    结果的排序算法，0~8, 分别是:
//											sort.Sort (go lib)，sort.Stable (go lib)，快速排序，堆排序，
//											归并排序，希尔排序，希尔排序(并发), 插入排序，选择排序
//			search_by	:FormValue int:    字符串搜索算法，0~3, 分别是:
//											regexp.FindAllIndex (go lib)，KMP 算法，Rabin-Karp 算法，暴力法
// Response:
//		Success: JSON: {"success", "token"}
//		Failed:  JSON: {"error": "error description"}
func (s *Service) apiWordfaPost(w http.ResponseWriter, r *http.Request) {
	// 获取 token, keywords, file, algorithm
	token := r.FormValue("token")

	keywords := r.FormValue("keywords")
	if keywords == "" {
		logging.Warning("apiWordfaPost failed: empty keywords")
		responseJson(&w, ErrorResponse{ErrorDescription: "Unexpected empty keywords"})
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		logging.Warning("apiWordfaPost failed: get FormFile Error:", err)
		responseJson(&w, ErrorResponse{ErrorDescription: "Cannot handle this file"})
		return
	}
	defer file.Close()

	sortAlgorithm, err := strconv.Atoi(r.FormValue("sort_by"))
	if err != nil || sortAlgorithm < 0 || sortAlgorithm > 8 {
		sortAlgorithm = sortalgo.StlSort
	}

	searchAlgorithm, err := strconv.Atoi(r.FormValue("search_by"))
	if err != nil || searchAlgorithm < 0 || searchAlgorithm > 3 {
		searchAlgorithm = strsearch.LibRe
	}

	// 停止该用户之前的任务
	if s, ok := s.WordFaSessionHolder.Get(token); ok {
		s.Task.Stop()
	}
	// 创建新任务
	task, err := s.buildTask(token, keywords, file, handler, searchAlgorithm)
	if err != nil {
		logging.Warning("apiWordfaPost failed: buildTask Error:", err)
		responseJson(&w, ErrorResponse{ErrorDescription: "Bad keywords or file given"})
		return
	}
	// 提交任务
	s.WordFaSessionHolder.Put(token, NewWordfaSession(task, sortAlgorithm))
	logging.Info(
		fmt.Sprintf("apiWordfaPost success: token=%#v\n\t--> file=%v\n\t--> keyword=%v\n\t--> sort by %v, search by %v",
			token, task.SrcFiles, task.Patterns, sortAlgorithm, searchAlgorithm,
		))
	responseJson(&w, PostApiWordfaResponse{Success: token})
	task.Run()
}

// buildTask 从请求解析出的数据构建一个 wordfa.Task
func (s *Service) buildTask(token string, keywords string, file multipart.File,
	handler *multipart.FileHeader, searchAlgorithm int) (*wordfa.Task, error) {

	defer file.Close()
	var task wordfa.Task
	// Patterns
	ks := strings.FieldsFunc(keywords, func(r rune) bool {
		return r == ',' || r == '，'
	})
	for i := 0; i < len(ks); i++ {
		if p := strings.TrimSpace(ks[i]); p != "" {
			task.Patterns = append(task.Patterns, p)
		}
	}

	// Algorithm
	task.StrSearchAlgorithm = searchAlgorithm

	// SrcFiles
	dir, fp, err := s.saveFile(token, file, handler)
	if err != nil {
		return &task, fmt.Errorf("system error: cannot create temp file: %s", err)
	}
	switch handler.Header.Get("Content-Type") {
	case "application/zip":
		if err = util.UnzipFile(dir, fp); err != nil {
			return &task, err
		}
		if task.SrcFiles, err = util.GetAllFiles(dir, "text/plain"); err != nil {
			return &task, err
		}
	case "text/plain":
		task.SrcFiles = []string{fp}
	default:
		return &task, fmt.Errorf("unsupported file type")
	}

	return &task, nil
}

// saveFile 在临时目录里保存请求的文件
func (s *Service) saveFile(token string, file multipart.File,
	handler *multipart.FileHeader) (dir string, fp string, err error) {

	dir, fp = s.tempFilePath(token, handler.Filename)
	// 创建临时目录
	if _, err := os.Stat(dir); err == nil {
		_ = os.RemoveAll(dir)
	}
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		logging.Debug(err)
		return "", "", err
	}
	// 创建、写入文件
	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logging.Debug(err)
		return "", "", err
	}
	defer f.Close()
	io.Copy(f, file)

	return dir, fp, nil
}

// tempFilePath 为请求文件获取临时目录名、文件名
func (s *Service) tempFilePath(token string, fileName string) (parentDir string, filePath string) {
	dirName := s.TempDirPrefix + base64.StdEncoding.EncodeToString([]byte(token))
	parentDir = filepath.FromSlash(path.Join(os.TempDir(), dirName))
	filePath = filepath.FromSlash(path.Join(parentDir, fileName))
	return parentDir, filePath
}
