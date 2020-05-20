// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package wordfa

import (
	"CiFa/util/sortalgo"
	"CiFa/util/strsearch"
	"io/ioutil"
	"sync"
)

// Task 是"统计给定关键词 Patterns 在一系列文本文件 SrcFiles 中出现的频数"的任务
//
// 可以通过对 StrSearchAlgorithm 字段赋值，以使用不同算法。
// 调用 Task 实例的 Run() 方法开始统计任务，
// 调用 Task 实例的 GetProgress() 方法获取任务执行进度，
// 在 Run 完成后，调用 Task 实例的 GetResult() 方法得到结果。
type Task struct {
	SrcFiles []string // 待检测的文件
	Patterns []string // 待匹配的词

	StrSearchAlgorithm int // 指定使用的字符串搜索算法, see strsearch

	fileMap map[string]bool // SrcFiles 中的所有文件，value 是代表是否检索完成的
	matches map[string]int  // 已完成的匹配 {"词": 出现次数}

	exit chan bool
	mux  sync.Mutex
}

func NewTask(srcFiles []string, patterns []string) *Task {
	return &Task{SrcFiles: srcFiles, Patterns: patterns}
}

// prepare before match
func (t *Task) prepare() {
	t.mux.Lock()
	defer t.mux.Unlock()

	// map patterns
	t.matches = map[string]int{}
	for _, p := range t.Patterns {
		t.matches[p] = 0
	}

	// Map files
	t.fileMap = map[string]bool{}
	for _, f := range t.SrcFiles {
		t.fileMap[f] = false
	}
}

// match search the files in Task.SrcFiles, try to get {"word": frequency} for each word in Task.Patterns
// Task.prepare() calling before this method is required
// Result put into Task.matches
func (t *Task) match() {
	if t.GetProgress() <= -1 {
		panic("Task not prepared, cannot run match()")
	}
	var wg sync.WaitGroup
	for filePath, _ := range t.fileMap {
		wg.Add(1)
		go func(t *Task, file string) {
			// Read File
			data, err := ioutil.ReadFile(file)
			if err != nil {
				panic(err)
			}
			// Find matches
			for _, pattern := range t.Patterns {
				found := len(strsearch.By(t.StrSearchAlgorithm).FindAllBytes(data, pattern))

				t.mux.Lock()
				t.matches[pattern] += found
				t.mux.Unlock()
			}
			// tag matched file
			t.mux.Lock()
			t.fileMap[file] = true
			t.mux.Unlock()

			wg.Done()
		}(t, filePath)
	}
	wg.Wait()
}

// GetProgress return the current status of Task
// Status:
//		<= -1		// unprepared
//		[0, 1)		// matching, the number means the progress
//		>= 1		// finished
func (t *Task) GetProgress() float32 {
	t.mux.Lock()
	defer t.mux.Unlock()

	// unprepared
	if t.fileMap == nil || t.matches == nil {
		return -1.1
	}

	totalFiles := 0
	finishedFiles := 0
	for _, finished := range t.fileMap {
		totalFiles++
		if finished {
			finishedFiles++
		}
	}

	// matching finished
	if finishedFiles == totalFiles {
		return 1.1
	}

	// matching, return progress
	return float32(finishedFiles) / float32(totalFiles)
}

// Run do preparing jobs and start matching task.
// It is Recommended to be called by:
//		go task.Run()
func (t *Task) Run() {
	// prepare
	if t.GetProgress() <= -1 {
		t.prepare()
	}

	// go run match
	t.mux.Lock()
	t.exit = make(chan bool)
	t.mux.Unlock()
	finished := make(chan bool)
	go func() {
		t.match()
		finished <- true
	}()

	// block until match finished or Stop called
	select {
	case <-finished:
		t.mux.Lock()
		t.exit = nil
		t.mux.Unlock()
		return
	case <-t.exit:
		t.mux.Lock()
		t.exit = nil
		t.mux.Unlock()
		return
	}
}

// Stop a Running Task
func (t *Task) Stop() {
	t.mux.Lock()
	defer t.mux.Unlock()

	if t.exit != nil {
		t.exit <- true
	}
}

// GetResult return the result matches (map[string]int) and ok=true if task is finished, (nil, false) else
func (t *Task) GetResult(sortAlgorithm int) (result Result, ok bool) {
	if t.GetProgress() >= 1 {
		t.mux.Lock()
		defer t.mux.Unlock()

		var result Result
		for k, f := range t.matches {
			result = append(result, ResultItem{
				Keyword:   k,
				Frequency: f,
			})
		}
		sortalgo.By(sortAlgorithm).Sort(result)
		return result, true
	}
	return nil, false
}

// Result 是 wordfa.Task 任务的结果，包含各给定关键词在文件中出现的频数
// Result 实现了 sort.Interface, 可以按频数从大到小排序
type Result []ResultItem

// ResultItem 是 Result 切片中的数据条目
type ResultItem struct {
	Keyword   string `json:"keyword"`
	Frequency int    `json:"frequency"`
}

func (r Result) Len() int {
	return len(r)
}

func (r Result) Less(i, j int) bool {
	return r[i].Frequency > r[j].Frequency
}

func (r Result) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
