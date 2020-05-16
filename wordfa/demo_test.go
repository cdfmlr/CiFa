// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

// THIS FILE CONTAINS A DEMO VERSION OF wordfa.
// IT IS NOW MEANINGLESS

package wordfa

//
//import (
//	"CiFa/util"
//	"CiFa/util/strsearch"
//	"fmt"
//	"io/ioutil"
//	"sync"
//	"testing"
//	"time"
//)
//
//func TestBaseDemo(t *testing.T) {
//	start := time.Now()
//
//	task := Task{
//		SrcDir:   "/Users/c/Desktop/glp",
//		Patterns: []string{"max", "name"},
//	}
//	//dir, err := ioutil.ReadDir(task.SrcDir)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//fmt.Printf("%#v\n", dir)
//	//for _, f := range dir {
//	//	fmt.Printf("\t%#v\n", f)
//	//}
//	fs, err := util.GetAllFiles(task.SrcDir, "text/plain")
//	if err != nil {
//		panic(err)
//	}
//	for _, f := range fs {
//		fmt.Printf("%#v\n", f)
//		data, err := ioutil.ReadFile(f)
//		if err != nil {
//			panic(err)
//		}
//		for _, p  := range task.Patterns {
//			fmt.Printf("\t%s: %v\n", p, strsearch.FindAllBytes(data, p))
//		}
//
//	}
//
//	fmt.Println(time.Since(start))
//}
//
//func TestConcurrentDemo(t *testing.T) {
//	// Expected:
//	// max: 29
//	// name: 14
//	start := time.Now()
//
//	task := Task{
//		SrcDir:   "/Users/c/Desktop/glp",
//		Patterns: []string{"max", "name"},
//	}
//	if err := task.prepare(); err != nil {
//		panic(err)
//	}
//	var wg sync.WaitGroup
//	for k, _ := range task.fileMap {
//		wg.Add(1)
//		go func (t *Task, file string) {
//			data, err := ioutil.ReadFile(file)
//			if err != nil {
//				panic(err)
//			}
//			for p, _  := range t.matches {
//				t.mux.Lock()
//				t.matches[p] += len(strsearch.FindAllBytes(data, p))
//				t.mux.Unlock()
//			}
//			t.mux.Lock()
//			defer t.mux.Unlock()
//			t.fileMap[file] = true
//			wg.Done()
//		}(&task, k)
//	}
//	wg.Wait()
//
//	task.mux.Lock()
//	fmt.Println(task.matches)
//	task.mux.Unlock()
//
//	fmt.Println(time.Since(start))
//
//	//
//	//finished := make(chan bool)
//	//go func(t *Task, finished chan bool) {
//	//	for {
//	//		unfinished := false
//	//		t.mux.Lock()
//	//		for _, v := range t.fileMap {
//	//			if !v {
//	//				unfinished = true
//	//				break
//	//			}
//	//		}
//	//		t.mux.Unlock()
//	//		if !unfinished {
//	//			finished <- true
//	//			return
//	//		}
//	//
//	//		time.Sleep(100 * time.Millisecond)
//	//	}
//	//}(&task, finished)
//	//
//	//select {
//	//case <- finished:
//	//	task.mux.Lock()
//	//	fmt.Println(task.matches)
//	//	task.mux.Unlock()
//	//}
//}
