// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package cliserve

import (
	"CiFa/util"
	"CiFa/util/sortalgo"
	"CiFa/wordfa"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type CliWordfaServer struct {
	KeywordFilePath string
	SourceFilePath  string

	SortAlgo      string
	StrsearchAlgo string

	OutputFilePath string
}

func (c *CliWordfaServer) Run() {
	patterns := getPatterns(c.KeywordFilePath)
	srcFiles := getSrcFiles(c.SourceFilePath)

	task := wordfa.NewTask(srcFiles, patterns)
	if c.SortAlgo != "" {
		task.SortFuncName = c.SortAlgo
	}
	if c.StrsearchAlgo != "" {
		task.StrSearchFuncName = c.StrsearchAlgo
	}

	//logging.Debug("patterns: ", task.Patterns)
	//logging.Debug("srcFiles: ", task.SrcFiles)

	go task.Run()

	finished := make(chan bool)
	go func() {
		for {
			p := task.GetProgress()
			if p < 0 {
				p = 0
			} else if p > 1 {
				p = 1
			}
			fmt.Printf("> progress: %.2f%%\n", p*100)
			if p >= 1 {
				finished <- true
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	if <-finished {
		if r, ok := task.GetResult(sortalgo.Heap); ok {
			if c.OutputFilePath != "" {
				if err := writeResultToFile(c.OutputFilePath, r); err == nil {
					fmt.Println("Result in", c.OutputFilePath)
					return
				} else {
					fmt.Println("Failed to write result: ", err)
				}
			}
			fmt.Println("Result: ")
			printResult(r)
			return
		}
	}
}

func printResult(result wordfa.Result) {
	for _, v := range result {
		fmt.Printf("%v: %v\n", v.Keyword, v.Frequency)
	}
}

func writeResultToFile(outFilePath string, result wordfa.Result) error {
	f, err := os.OpenFile(
		outFilePath,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0600,
	)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, v := range result {
		f.Write([]byte(fmt.Sprintf("%v: %v\n", v.Keyword, v.Frequency)))
	}
	return nil
}

// getPatterns 从文件 keywordFilePath 里获取要匹配的子串
func getPatterns(keywordFilePath string) []string {
	patterns := make([]string, 0)

	data, err := ioutil.ReadFile(keywordFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	ks := strings.FieldsFunc(string(data), func(r rune) bool {
		return r == ',' || r == '，' || r == '\n' || r == '\r'
	})
	for i := 0; i < len(ks); i++ {
		if p := strings.TrimSpace(ks[i]); p != "" {
			patterns = append(patterns, p)
		}
	}
	return patterns
}

// getSrcFiles 从文件/目录 sourceFilePath 里获取要匹配的文件
// 若 sourceFilePath 是目录则递归寻找其中所有类型为 text/plain 的文件
// 若 sourceFilePath 单个文件则返回[]string{sourceFilePath}
func getSrcFiles(sourceFilePath string) []string {
	srcFiles := make([]string, 0)

	f, err := os.Open(sourceFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	s, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	if s.IsDir() {
		if srcFiles, err = util.GetAllFiles(sourceFilePath, "text/plain"); err != nil {
			log.Fatalln(err)
		}
	} else {
		srcFiles = []string{sourceFilePath}
		//ft := util.GetFileContentType(f)
		//logging.Debug(ft)
		//if len(strsearch.FindAll(ft, "text/plain")) <= 0 {
		//	log.Fatalln("cannot handle a single file whose type is not text/plain")
		//} else {
		//	srcFiles = append(srcFiles, sourceFilePath)
		//}
	}
	return srcFiles
}
