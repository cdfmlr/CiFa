// Copyright (c) 2020 CDFMLR. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at  http://www.apache.org/licenses/LICENSE-2.0

package sortalgo

import (
	"sort"
	"sync"
)

// selection sort
func SelectionSort(data sort.Interface, a, b int) {
	for i := a; i <= b; i++ {
		smallest := i
		for j := i + 1; j <= b; j++ {
			if data.Less(j, smallest) {
				smallest = j
			}
		}
		data.Swap(i, smallest)
	}
}

// insertion sort
func InsertionSort(data sort.Interface, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

// shell sort
func ShellSort(data sort.Interface, a, b int) {
	for step := int(uint(b-a+1) >> 1); step > 0; step >>= 1 {
		for i := a + step; i <= b; i++ {
			for j := i - step; j >= a && data.Less(j+step, j); j -= step {
				data.Swap(j, j+step)
			}
		}
	}
}

// <abandoned> shell sort 的并发实现，效率不如 `ShellSort`，不推荐使用
func ShellSortSync(data sort.Interface, a, b int) {
	var wg sync.WaitGroup
	for step := int(uint(b-a+1) >> 1); step > 0; step >>= 1 {
		for i := a + step; i <= b; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				for j := i - step; j >= a && data.Less(j+step, j); j -= step {
					data.Swap(j, j+step)
				}
			}(i)
		}
		wg.Wait()
	}
}

// quick sort
func QuickSort(data sort.Interface, a, b int) {
	if a < b {
		m := partition(data, a, b)
		QuickSort(data, a, m-1)
		QuickSort(data, m+1, b)
	}
}

// partition-exchange for quick sort
func partition(data sort.Interface, a, b int) (m int) {
	m = a
	for u := a; u < b; u++ {
		if data.Less(u, b) {
			data.Swap(m, u)
			m += 1
		}
	}
	data.Swap(m, b)
	return m
}

// heap sort
func HeapSort(data sort.Interface, a, b int) {
	heapSize := b - a
	buildMaxHeap(data, a, b, heapSize)
	for i := b; i > 0; i-- {
		data.Swap(0, i)
		heapSize--
		maxHeapify(data, 0, heapSize)
	}
}

func buildMaxHeap(data sort.Interface, a, b int, heapSize int) {
	for i := b / 2; i >= a; i-- {
		maxHeapify(data, i, heapSize)
	}
}

// maxHeapify maintain the max-heap property.
func maxHeapify(data sort.Interface, i int, heapSize int) {
	l := i << 1   // left
	r := 1 + i<<1 // right
	var largest int
	if l <= heapSize && data.Less(i, l) {
		largest = l
	} else {
		largest = i
	}
	if r <= heapSize && data.Less(largest, r) {
		largest = r
	}
	if largest != i {
		data.Swap(i, largest)
		maxHeapify(data, largest, heapSize)
	}
}

// merge sort
func MergeSort(data sort.Interface, a, b int) {
	if a < b {
		mid := int(uint(a+b) >> 1)
		MergeSort(data, a, mid)
		MergeSort(data, mid+1, b)
		symMerge(data, a, mid+1, b+1)
		// 这里的 mid+1，b+1 是为了满足 SymMerge，
		// SymMerge 是对 data[a:m] and data[m:b] 归并，
		// 末尾的 m、b 要 +1 才是操作了所有数据
	}
}

// SymMerge merges the two sorted subsequences data[a:m] and data[m:b] using
// the SymMerge algorithm from Pok-Son Kim and Arne Kutzner, "Stable Minimum
// Storage Merging by Symmetric Comparisons", in Susanne Albers and Tomasz
// Radzik, editors, Algorithms - ESA 2004, volume 3221 of Lecture Notes in
// Computer Science, pages 714-723. Springer, 2004.
//
// Let M = m-a and N = b-n. Wolog M < N.
// The recursion depth is bound by ceil(log(N+M)).
// The algorithm needs O(M*log(N/M + 1)) calls to data.Less.
// The algorithm needs O((M+N)*log(M)) calls to data.Swap.
//
// The paper gives O((M+N)*log(M)) as the number of assignments assuming a
// rotation algorithm which uses O(M+N+gcd(M+N)) assignments. The argumentation
// in the paper carries through for Swap operations, especially as the block
// swapping rotate uses only O(M+N) Swaps.
//
// symMerge assumes non-degenerate arguments: a < m && m < b.
// Having the caller check this condition eliminates many leaf recursion calls,
// which improves performance.
//
// Copy From go lib src：`sort/sort.go`
// Copyright 2009 The Go Authors.
func symMerge(data sort.Interface, a, m, b int) {
	// 首先，基本情况：当从左/右端到中间只有一个元素时
	// 直接把它插入到另一个子数组中的正确排序位置

	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[a] into data[m:b]
	// if data[a:m] only contains one element.
	if m-a == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] >= data[a] for m <= i < b.
		// Exit the search loop with i == b in case no such index exists.
		i := m
		j := b
		for i < j {
			h := int(uint(i+j) >> 1)
			if data.Less(h, a) {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[a] reaches the position before i.
		for k := a; k < i-1; k++ {
			data.Swap(k, k+1)
		}
		return
	}

	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[m] into data[a:m]
	// if data[m:b] only contains one element.
	if b-m == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] > data[m] for a <= i < m.
		// Exit the search loop with i == m in case no such index exists.
		i := a
		j := m
		for i < j {
			h := int(uint(i+j) >> 1)
			if !data.Less(m, h) {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[m] reaches the position i.
		for k := m; k > i; k-- {
			data.Swap(k, k-1)
		}
		return
	}

	// 找出合适的区间以进行旋转

	mid := int(uint(a+b) >> 1)
	n := mid + m
	var start, r int
	if m > mid {
		start = n - b
		r = mid
	} else {
		start = a
		r = m
	}
	p := n - 1

	for start < r {
		c := int(uint(start+r) >> 1)
		if !data.Less(p-c, c) {
			start = c + 1
		} else {
			r = c
		}
	}

	// 对区间进行旋转，然后递归合并子切片

	end := n - start
	if start < m && m < end {
		rotate(data, start, m, end)
	}
	if a < start && start < mid {
		symMerge(data, a, start, mid)
	}
	if mid < end && end < b {
		symMerge(data, mid, end, b)
	}
}

// Rotate two consecutive blocks u = data[a:m] and v = data[m:b] in data:
// Data of the form 'x u v y' is changed to 'x v u y'.
// Rotate performs at most b-a many calls to data.Swap.
// Rotate assumes non-degenerate arguments: a < m && m < b.
//
// Copy From go lib src：`sort/sort.go`
// Copyright 2009 The Go Authors.
func rotate(data sort.Interface, a, m, b int) {
	i := m - a
	j := b - m

	for i != j {
		if i > j {
			swapRange(data, m-i, m, j)
			i -= j
		} else {
			swapRange(data, m-i, m+j-i, i)
			j -= i
		}
	}
	// i == j
	swapRange(data, m-i, m, i)
}

// swapRange
// Copy From go lib src：`sort/sort.go`
// Copyright 2009 The Go Authors.
func swapRange(data sort.Interface, a, b, n int) {
	for i := 0; i < n; i++ {
		data.Swap(a+i, b+i)
	}
}
