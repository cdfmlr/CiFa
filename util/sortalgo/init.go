// sortalgo 包提供各种排序算法
//
// Algorithms Implemented:
//		- SelectionSort
//		- InsertionSort
//		- ShellSort
//		- ShellSortSync (abandoned)
//		- QuickSort
//		- HeapSort
//
// User Interface:
//	1. sortalgo.Sort
//			sortalgo.Sort(data)
//		A shorthand to `sortalgo.By(sortalgo.StlSort).Sort(data)`
//	2. sortalgo.By
//			sortalgo.By(ALGORITHM).Sort(data)
//		Sort data by ALGORITHM where ALGORITHM can be:
//			Quick		// 快速排序
//			Heap		// 堆排序
//			Merge		// 归并排序
//			Shell		// 希尔排序
//			ShellSync	// 希尔排序(并发), 不推荐
//			Insertion	// 插入排序
//			Selection	// 选择排序
//			StlSort		// sort.Sort (go lab)
//			StlStable	// sort.Stable (go lab)
// 	Example:
//		type dataIntS []int
//
//		func (d dataIntS) Len() int {
//			return len(d)
//		}
//
//		func (d dataIntS) Less(i, j int) bool {
//			return d[i] < d[j]
//		}
//
//		func (d dataIntS) Swap(i, j int) {
//			d[i], d[j] = d[j], d[i]
//		}
//
//		func main () {
//			data := dataIntS{2, 1, 3, 7, 4}
//			sortalgo.By(sortalgo.Quick).Sort(data)
//			fmt.Println(data)	// [1, 2, 3, 4, 7]
//		}
//

package sortalgo

import (
	"sort"
)

// Algorithms
const (
	StlSort   = iota //"sort.Sort (go lab)"
	StlStable        //"sort.Stable (go lab)"
	Quick            //"快速排序"
	Heap             //"堆排序"
	Merge            //"归并排序"
	Shell            //"希尔排序"
	ShellSync        //"希尔排序(并发), 不推荐"
	Insertion        //"插入排序"
	Selection        //"选择排序"
	_nothing
)

type SortAlgo int

func By(algorithm int) SortAlgo {
	if algorithm < 0 || algorithm >= _nothing {
		panic("Unknown algorithm")
	}
	return SortAlgo(algorithm)
}

func (s SortAlgo) Sort(data sort.Interface) {
	var fun func(p sort.Interface, a, b int)
	switch s {
	case StlSort:
		sort.Sort(data)
		return
	case StlStable:
		sort.Stable(data)
		return
	case Quick:
		fun = QuickSort
	case Heap:
		fun = HeapSort
	case Merge:
		fun = MergeSort
	case Shell:
		fun = ShellSort
	case ShellSync:
		fun = ShellSortSync
	case Insertion:
		fun = InsertionSort
	case Selection:
		fun = SelectionSort
	}
	if fun != nil {
		fun(data, 0, data.Len()-1)
	} else {
		panic("Unknown algorithm")
	}
}

func Sort(data sort.Interface) {
	By(StlSort).Sort(data)
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
