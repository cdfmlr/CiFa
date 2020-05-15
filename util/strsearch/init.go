// strsearch 包提供一系列字符串搜索算法
//
// algorithms implemented:
//  - NaiveSearchByChar
//  - NaiveSearchBySlice
//  - KmpSearch
//  - RabinKarpSearch
//
// Notes:
//  NaiveSearchBySlice is slower than NaiveSearchByChar
//  RabinKarpSearch is despised, for its md5 calling as a hash function, it's too slowwwwww.
//
// Usage:
// 		strsearch.By(strsearch.ALGORITHM).FindAll/FindAllBytes(text, pattern)
// 	ALGORITHM may be:
//		Naive		// 暴力法
//		Kmp			// KMP 算法
//		RabinKarp	// Rabin-Karp 算法
//		LibRe		// regexp.FindAllIndex (go lib)

package strsearch

import "regexp"

//Algorithms
const (
	Naive     = iota // 暴力法
	Kmp              // KMP 算法
	RabinKarp        // Rabin-Karp 算法
	LibRe            // regexp.FindAllIndex (go lib)
	_nothing
)

type StrSearch int

func By(algorithm int) StrSearch {
	if algorithm < 0 || algorithm >= _nothing {
		panic("Unknown algorithm")
	}
	return StrSearch(algorithm)
}

// FindAll 在 text (string) 中搜索 pattern，返回所有匹配位置起点索引
func (s StrSearch) FindAll(text string, pattern string) []int {
	switch s {
	case Naive:
		return NaiveSearchByChar(text, pattern, -1)
	case Kmp:
		return KmpSearch(text, pattern, -1)
	case RabinKarp:
		return RabinKarpSearch(text, pattern, -1)
	case LibRe:
		return RegSearch(text, pattern, -1)
	}
	panic("Unknown algorithm")
}

// FindAllBytes 在 text ([]byte) 中搜索 pattern，返回所有匹配位置起点索引
// 若 StrSearch 为 LibRe，则会使用 RegSearchBytes 完成检索；
// 否则会返回 s.FindAll(string(text), pattern) 的调用结果。
func (s StrSearch) FindAllBytes(text []byte, pattern string) []int {
	if s == LibRe {
		return RegSearchBytes(text, pattern, -1)
	}
	return s.FindAll(string(text), pattern)
}

func FindAll(text string, pattern string) []int {
	return By(LibRe).FindAll(text, pattern)
}

func FindAllBytes(text []byte, pattern string) []int {
	return By(LibRe).FindAllBytes(text, pattern)
}

func RegSearch(s, substr string, maxMatches int) (indices []int) {
	reg := regexp.MustCompile(substr)
	r := reg.FindAllStringIndex(s, maxMatches)
	var res []int
	for _, i := range r {
		res = append(res, i[0])
	}
	return res
}

func RegSearchBytes(b []byte, pattern string, maxMatches int) (indices []int) {
	reg := regexp.MustCompile(pattern)
	r := reg.FindAllIndex(b, maxMatches)
	var res []int
	for _, i := range r {
		res = append(res, i[0])
	}
	return res
}

/******************************************************************************
 *    Copyright 2020 CDFMLR                                                   *
 *                                                                            *
 *    Licensed under the Apache License, Version 2.0 (the "License");         *
 *    you may not use this file except in compliance with the License.        *
 *    You may obtain a copy of the License at                                 *
 *                                                                            *
 *        http://www.apache.org/licenses/LICENSE-2.0                          *
 *                                                                            *
 *    Unless required by applicable law or agreed to in writing, software     *
 *    distributed under the License is distributed on an "AS IS" BASIS,       *
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.*
 *    See the License for the specific language governing permissions and     *
 *    limitations under the License.                                          *
 *                                                                            *
 ******************************************************************************/
