// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package util

import (
	"CiFa/util/strsearch"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// GetAllFiles 获取 dirPth 目录下的所有文件, 包含子目录下的文件
// 可以对 spicType 参数传递一个 MIME Content-Type 字符串（e.g. "text/plain"），以只获取指定类型的文件
// spicType 传入空字符串("")获取任意类型的文件
func GetAllFiles(dirPth string, spicType string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return files, err
	}
	for _, f := range dir {
		p := filepath.FromSlash(path.Join(dirPth, f.Name()))

		if f.IsDir() {
			fs, err := GetAllFiles(p, spicType)
			files = append(files, fs...)
			if err != nil {
				return files, err
			}
			continue
		} else if spicType != "" {
			file, err := os.Open(p)
			if err != nil {
				if file != nil {
					_ = file.Close()
				}
				return files, err
			}
			fileType := GetFileContentType(file)
			_ = file.Close()
			if len(strsearch.FindAll(fileType, spicType)) <= 0 {
				continue
			}
		}
		files = append(files, p)
	}
	return files, err
}

// GetFileContentType 获取一个文件的 MIME Content-Type
func GetFileContentType(file *os.File) string {
	buffer := make([]byte, 512) // sniffLen = 512
	_, err := file.Read(buffer)
	if err != nil {
		return "text/plain; charset=utf-8"
	}
	return http.DetectContentType(buffer)
}
