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

//Algorithms
const (
	LibRe     = iota // regexp.FindAllIndex (go lib)
	Kmp              // KMP 算法
	RabinKarp        // Rabin-Karp 算法
	Naive            // 暴力法
	_nothing
)

func By(algorithm int) StrSearchAlgorithm {
	if algorithm < 0 || algorithm >= _nothing {
		panic("Unknown algorithm")
	}
	var strSearchAlgo StrSearchAlgorithm
	switch algorithm {
	case LibRe:
		strSearchAlgo = goStlRegSearch
	case Kmp:
		strSearchAlgo = KmpSearch
	case RabinKarp:
		strSearchAlgo = RabinKarpSearch
	case Naive:
		strSearchAlgo = NaiveSearchByChar
	}
	return strSearchAlgo
}

func FindAll(text string, pattern string) []int {
	return By(LibRe).FindAll(text, pattern)
}

func FindAllBytes(text []byte, pattern string) []int {
	return By(LibRe).FindAllBytes(text, pattern)
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
