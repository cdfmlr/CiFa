// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package sortalgo

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

// Test data struct: []int
type dataIntS []int

func (d dataIntS) Len() int {
	return len(d)
}

func (d dataIntS) Less(i, j int) bool {
	return d[i] < d[j]
}

func (d dataIntS) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Test data struct: []string
type dataStrS []string

func (d dataStrS) Len() int {
	return len(d)
}

func (d dataStrS) Less(i, j int) bool {
	return d[i] > d[j]
}

func (d dataStrS) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Test data struct: []struct by a field
type dataPersonSByAge []struct {
	Name string
	Age  int
}

func (p dataPersonSByAge) Len() int {
	return len(p)
}

func (p dataPersonSByAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p dataPersonSByAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// common tests field for all sort algorithms
type args struct {
	data sort.Interface
	a    int
	b    int
}

var tests = []struct {
	name   string
	args   args
	expect sort.Interface
}{
	{
		name: "ints",
		args: args{
			data: dataIntS{2, 1, 3, 6, 7, 9, 8, 5, 4},
			a:    0,
			b:    8,
		},
		expect: dataIntS{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		name: "strs_reverse",
		args: args{
			data: dataStrS{"qde", "abc", "xyz", "zab", "cdd", "axz"},
			a:    0,
			b:    5,
		},
		expect: dataStrS{"zab", "xyz", "qde", "cdd", "axz", "abc"},
	},
	{
		name: "structs_part",
		args: args{
			data: dataPersonSByAge{{
				Name: "Foo",
				Age:  34,
			}, {

				Name: "Fuzz",
				Age:  12,
			}, {
				Name: "Bar",
				Age:  4,
			}, {
				Name: "Buzz",
				Age:  1,
			}, {
				Name: "John Doe",
				Age:  66,
			}, {
				Name: "Jane Doe",
				Age:  8,
			}},
			a: 0,
			b: 3,
		},
		expect: dataPersonSByAge{{
			Name: "Buzz",
			Age:  1,
		}, {
			Name: "Bar",
			Age:  4,
		}, {
			Name: "Fuzz",
			Age:  12,
		}, {
			Name: "Foo",
			Age:  34,
		}, {
			Name: "John Doe",
			Age:  66,
		}, {
			Name: "Jane Doe",
			Age:  8,
		}},
	},
}

func TestHeapSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.data, tt.args.a, tt.args.b)
			if !reflect.DeepEqual(tt.args.data, tt.expect) {
				t.Errorf("tt.args.data != tt.expect:\n\t-->result:%v\n\t-->expect:%v", tt.args.data, tt.expect)
			} else {
				t.Log("Success sorted:", tt.args.data)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.data, tt.args.a, tt.args.b)
			if !reflect.DeepEqual(tt.args.data, tt.expect) {
				t.Errorf("tt.args.data != tt.expect:\n\t-->result:%v\n\t-->expect:%v", tt.args.data, tt.expect)
			} else {
				t.Log("Success sorted:", tt.args.data)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.data, tt.args.a, tt.args.b)
			if !reflect.DeepEqual(tt.args.data, tt.expect) {
				t.Errorf("tt.args.data != tt.expect:\n\t-->result:%v\n\t-->expect:%v", tt.args.data, tt.expect)
			} else {
				t.Log("Success sorted:", tt.args.data)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.data, tt.args.a, tt.args.b)
			if !reflect.DeepEqual(tt.args.data, tt.expect) {
				t.Errorf("tt.args.data != tt.expect:\n\t-->result:%v\n\t-->expect:%v", tt.args.data, tt.expect)
			} else {
				t.Log("Success sorted:", tt.args.data)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.args.data, tt.args.a, tt.args.b)
			if !reflect.DeepEqual(tt.args.data, tt.expect) {
				t.Errorf("tt.args.data != tt.expect:\n\t-->result:%v\n\t-->expect:%v", tt.args.data, tt.expect)
			} else {
				t.Log("Success sorted:", tt.args.data)
			}
		})
	}
}

func TestShellSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShellSort(tt.args.data, tt.args.a, tt.args.b)
			if !reflect.DeepEqual(tt.args.data, tt.expect) {
				t.Errorf("tt.args.data != tt.expect:\n\t-->result:%v\n\t-->expect:%v", tt.args.data, tt.expect)
			} else {
				t.Log("Success sorted:", tt.args.data)
			}
		})
	}
}

func TestShellSortSync(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShellSortSync(tt.args.data, tt.args.a, tt.args.b)
			if !reflect.DeepEqual(tt.args.data, tt.expect) {
				t.Errorf("tt.args.data != tt.expect:\n\t-->result:%v\n\t-->expect:%v", tt.args.data, tt.expect)
			} else {
				t.Log("Success sorted:", tt.args.data)
			}
		})
	}
}

func TestEff(t *testing.T) {
	// random test data
	rand.Seed(time.Now().UnixNano())
	n := 100000000
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(n))
	}
	var data dataIntS = arr

	var elapsed time.Duration

	elapsed = sortIntSElapsedJudge(HeapSort, data, 0, n-1)
	fmt.Println("HeapSort:\t\t", elapsed)

	elapsed = sortIntSElapsedJudge(MergeSort, data, 0, n-1)
	fmt.Println("MergeSort:\t\t", elapsed)

	elapsed = sortIntSElapsedJudge(QuickSort, data, 0, n-1)
	fmt.Println("QuickSort:\t\t", elapsed)

	elapsed = sortIntSElapsedJudge(ShellSort, data, 0, n-1)
	fmt.Println("ShellSort:\t\t", elapsed)

	//elapsed = sortIntSElapsedJudge(ShellSortSync, data, 0, n-1)
	//fmt.Println("ShellSortSync:\t", elapsed)

	//elapsed = sortIntSElapsedJudge(InsertionSort, data, 0, n-1)
	//fmt.Println("InsertionSort:\t", elapsed)
	//
	//elapsed = sortIntSElapsedJudge(SelectionSort, data, 0, n-1)
	//fmt.Println("SelectionSort:\t", elapsed)
}

func sortIntSElapsedJudge(sortF func(data sort.Interface, a, b int), data sort.Interface, a, b int) (elapsed time.Duration) {
	//fmt.Println(data)
	var dataCopy = dataIntS{}
	for _, i := range data.(dataIntS) {
		dataCopy = append(dataCopy, i)
	}
	//reflect.Copy(reflect.ValueOf(dataCopy), reflect.ValueOf(data))
	t := time.Now()
	sortF(dataCopy, a, b)
	elapsed = time.Since(t)
	//fmt.Println(data, dataCopy)
	return elapsed
}
