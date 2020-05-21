// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package strsearch

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

type args struct {
	s          string
	substr     string
	maxMatches int
}

var tests = []struct {
	name        string
	args        args
	wantIndices []int
}{
	{
		name: "abc-d.all",
		args: args{
			s:          "abc",
			substr:     "d",
			maxMatches: -1,
		},
		wantIndices: nil,
	},
	{
		name: "abc-aa.all",
		args: args{
			s:          "abc",
			substr:     "aa",
			maxMatches: -1,
		},
		wantIndices: nil,
	},
	{
		name: "abcb-b.all",
		args: args{
			s:          "aabbccbb",
			substr:     "bb",
			maxMatches: -1,
		},
		wantIndices: []int{2, 6},
	},
	{
		name: "abcb-bb.all",
		args: args{
			s:          "aabbccbb",
			substr:     "bbb",
			maxMatches: -1,
		},
		wantIndices: nil,
	},
	{
		name: "abcb-b.1",
		args: args{
			s:          "aabbccbb",
			substr:     "bb",
			maxMatches: 1,
		},
		wantIndices: []int{2},
	},
	{
		name: "nil-a.all",
		args: args{
			s:          "",
			substr:     "a",
			maxMatches: -1,
		},
		wantIndices: nil,
	},
	{
		name: "a-nil.all",
		args: args{
			s:          "s",
			substr:     "",
			maxMatches: -1,
		},
		wantIndices: nil,
	},
	{
		name: "text-single-match",
		args: args{
			s:          "We now present a linear-time string-matching algorithm due to Knuth, Morris, and Pratt.",
			substr:     "algorithm",
			maxMatches: 0,
		},
		wantIndices: []int{45},
	},
	{
		name: "chinese-text",
		args: args{
			s:          "神经网络是一个个「层」组成的。一个「层」就像是一个“蒸馏过滤器”，它会“过滤”处理输入的数据，从里面“精炼”出需要的信息，然后传到下一层。这样一系列的「层」组合起来，像流水线一样对数据进行处理。层层扬弃，让被处理的数据，或者说“数据的表示”对我们最终希望的结果越来越“有用”。",
			substr:     "层",
			maxMatches: 0,
		},
		wantIndices: []int{27, 54, 201, 228, 291, 294},
	},
	{
		name: "long-text",
		args: args{
			s: `Before there were computers, there were algorithms. But now that there are computers, there are even more algorithms, and algorithms lie at the heart of computing.
This book provides a comprehensive introduction to the modern study of computer algorithms. It presents many algorithms and covers them in considerable depth, yet makes their design and analysis accessible to all levels of readers. We have tried to keep explanations elementary without sacrificing depth of coverage or mathematical rigor.
Each chapter presents an algorithm, a design technique, an application area, or a related topic. Algorithms are described in English and in a pseudocode designed to be readable by anyone who has done a little programming. The book contains 244 figures—many with multiple parts—illustrating how the algorithms work. Since we emphasize efficiency as a design criterion, we include careful analyses of the running times of all our algorithms.
The text is intended primarily for use in undergraduate or graduate courses in algorithms or data structures. Because it discusses engineering issues in algorithm design, as well as mathematical aspects, it is equally well suited for self-study by technical professionals.
In this, the third edition, we have once again updated the entire book. The changes cover a broad spectrum, including new chapters, revised pseudocode, and a more active writing style.`,
			substr:     "algorithm",
			maxMatches: 0,
		},
		wantIndices: []int{40, 106, 122, 244, 273, 528, 805, 935, 1026, 1100},
	},
}

func TestNaiveSearchByChar(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Log(NaiveSearchByChar(tt.args.s, tt.args.substr, tt.args.maxMatches))
			if gotIndices := NaiveSearchByChar(tt.args.s, tt.args.substr, tt.args.maxMatches); !reflect.DeepEqual(gotIndices, tt.wantIndices) {
				t.Errorf("NaiveSearchByChar() = %v, want %v", gotIndices, tt.wantIndices)
			}
		})
	}
}

func TestNaiveSearchBySlice(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndices := NaiveSearchBySlice(tt.args.s, tt.args.substr, tt.args.maxMatches); !reflect.DeepEqual(gotIndices, tt.wantIndices) {
				t.Errorf("NaiveSearchBySlice() = %v, want %v", gotIndices, tt.wantIndices)
			}
		})
	}
}

func TestKmpSearch(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndices := KmpSearch(tt.args.s, tt.args.substr, tt.args.maxMatches); !reflect.DeepEqual(gotIndices, tt.wantIndices) {
				t.Errorf("KmpSearch() = %v, want %v", gotIndices, tt.wantIndices)
			}
		})
	}
}

func TestRabinKarpSearch(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndices := RabinKarpSearch(tt.args.s, tt.args.substr, tt.args.maxMatches); !reflect.DeepEqual(gotIndices, tt.wantIndices) {
				t.Errorf("RabinKarpSearch() = %v, want %v", gotIndices, tt.wantIndices)
			}
		})
	}
}

func TestEff(t *testing.T) {
	data, err := ioutil.ReadFile("testing_text.txt")
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	//pattern :=  "那么，"
	//pattern := "一九一八年"
	pattern := "阿Ｑ要画圆圈了，那手捏着笔却只是抖。于是那人替他将纸铺在地上，阿Ｑ伏下去，使尽了平生的力气画圆圈。他生怕被人笑话，立志要画得圆，但这可恶的笔不但很沉重，并且不听话，刚刚一抖一抖的几乎要合缝，却又向外一耸，画成瓜子模样了。"

	var elapsed time.Duration
	var res []int

	elapsed, res = stringMatchElapsedJudge(NaiveSearchByChar, text, pattern, 0)
	fmt.Println("NaiveSearchByChar:\t", elapsed, res)

	elapsed, res = stringMatchElapsedJudge(NaiveSearchBySlice, text, pattern, 0)
	fmt.Println("NaiveSearchBySlice:\t", elapsed, res)

	elapsed, res = stringMatchElapsedJudge(KmpSearch, text, pattern, 0)
	fmt.Println("KmpSearch:\t\t\t", elapsed, res)

	elapsed, res = stringMatchElapsedJudge(RabinKarpSearch, text, pattern, 0)
	fmt.Println("RabinKarpSearch:\t", elapsed, res)

	now := time.Now()
	//reg := regexp.MustCompile(pattern)
	//r := reg.FindAllIndex(data, -1)
	res = goStlRegSearch(text, pattern, -1)
	elapsed = time.Since(now)
	fmt.Println("regexp FindAllStringIndex:", elapsed, res)

	now = time.Now()
	res = goStlRegSearchBytes(data, pattern, -1)
	elapsed = time.Since(now)
	fmt.Println("regexp FindAllIndex:", elapsed, res)
}

func stringMatchElapsedJudge(algorithm func(s, substr string, maxMatches int) (indices []int),
	s, substr string, maxMatches int) (elapsed time.Duration, res []int) {

	t := time.Now()
	res = algorithm(s, substr, maxMatches)
	elapsed = time.Since(t)
	return elapsed, res
}

func TestInterface(t *testing.T) {
	data, err := ioutil.ReadFile("testing_text.txt")
	if err != nil {
		t.Fatal(err)
	}
	pattern := "一九一八年"
	r := By(LibRe).FindAllBytes(data, pattern)
	t.Log(r)
}
