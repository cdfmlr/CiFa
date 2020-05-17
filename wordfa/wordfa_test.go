// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package wordfa

import (
	"CiFa/util"
	"CiFa/util/sortalgo"
	"CiFa/util/strsearch"
	"fmt"
	"testing"
	"time"
)

func TestWordfaTask(t *testing.T) {
	files, _ := util.GetAllFiles("/Users/c/Desktop/glp", "text/plain")

	//task := Task{
	//	SrcFiles: files,
	//	Patterns: []string{"max", "name"},
	//}

	task := NewTask(files, []string{"opt", "max", "没有的东西", "name"})
	task.StrSearchAlgorithm = strsearch.Kmp
	go task.Run()

	finished := make(chan bool)
	go func() {
		for {
			p := task.GetProgress()
			fmt.Println("----> progress: ", p)
			if p >= 1 {
				finished <- true
				return
			}
			time.Sleep(10 * time.Microsecond)
		}
	}()

	if <-finished {
		if r, ok := task.GetResult(sortalgo.Heap); ok {
			t.Log("Result:", r)
			return
		}
	}
	t.Error("Failed")
}
