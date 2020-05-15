// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package strsearch

import (
	"crypto/md5"
	"encoding/binary"
	"io"
)

// naive string search algorithm
func NaiveSearchByChar(s, substr string, maxMatches int) (indices []int) {
	if len(s) == 0 || len(substr) == 0 || len(substr) > len(s) {
		return indices
	}
	for i := 0; i < len(s)-len(substr)+1; i++ {
		for j := 0; j < len(substr); j++ {
			if s[i+j] != substr[j] {
				break
			}
			if j == len(substr)-1 {
				indices = append(indices, i)
				if maxMatches > 0 && len(indices) >= maxMatches {
					return indices
				}
			}
		}
	}
	return indices
}

func NaiveSearchBySlice(s, substr string, maxMatches int) (indices []int) {
	if len(s) == 0 || len(substr) == 0 || len(substr) > len(s) {
		return indices
	}
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if s[i:i+len(substr)] == substr {
			indices = append(indices, i)
			if maxMatches > 0 && len(indices) >= maxMatches {
				return indices
			}
		}
	}
	return indices
}

// KMP algorithm
func KmpSearch(s, substr string, maxMatches int) (indices []int) {
	if len(s) == 0 || len(substr) == 0 || len(substr) > len(s) {
		return indices
	}
	next := computePrefixFunction(substr)
	numMatchedChar := -1
	for i := 0; i < len(s); i++ {
		for numMatchedChar > -1 && substr[numMatchedChar+1] != s[i] {
			numMatchedChar = next[numMatchedChar]
		}
		if substr[numMatchedChar+1] == s[i] {
			numMatchedChar++
		}
		if numMatchedChar == len(substr)-1 {
			indices = append(indices, i-len(substr)+1)
			if maxMatches > 0 && len(indices) >= maxMatches {
				return indices
			}
			numMatchedChar = next[numMatchedChar]
		}
	}
	return indices
}

func computePrefixFunction(substr string) []int {
	pi := make([]int, len(substr))
	pi[0] = -1
	k := 0
	for i := 1; i < len(substr); i++ {
		for k > 0 && substr[k+1] != substr[i] {
			k = pi[k]
		}
		if substr[k+1] == substr[i] {
			k++
		}
		pi[i] = k - 1
	}
	return pi
}

// Rabin-Karp algorithm
func RabinKarpSearch(s, substr string, maxMatches int) (indices []int) {
	if len(s) == 0 || len(substr) == 0 || len(substr) > len(s) {
		return indices
	}
	hsubstr := hash(substr)
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if hs := hash(s[i : i+len(substr)]); hs == hsubstr {
			if s[i:i+len(substr)] == substr { // s[i:i+len(substr)] == substr
				indices = append(indices, i)
				if maxMatches > 0 && len(indices) >= maxMatches {
					return indices
				}
			}
		}
	}
	return indices
}

func hash(s string) uint64 {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	b := h.Sum(nil)
	return binary.LittleEndian.Uint64(b)
}
