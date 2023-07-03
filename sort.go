package alglib

import (
	"alglib/generics"
)

func IsSorted[N generics.Sortable](list []N) bool {
	return IsSortedWithComparator(list, generics.ComparatorLesserEqual[N])
}

func IsSortedWithComparator[T any](list []T, le_comparator func(T, T) bool) bool {
	for i := 1; i < len(list); i++ {
		if !le_comparator(list[i-1], list[i]) {
			return false
		}
	}

	return true
}

func InsertionSort[N generics.Sortable](list []N) []N {
	return InsertionSortWithComparator(list, generics.ComparatorLessThan[N])
}

func InsertionSortWithComparator[T any](list []T, lt_comparator func(T, T) bool) []T {

	for indexToSortAgainst, _ := range list {
		indexToSwap := indexToSortAgainst
		valueToSwap := list[indexToSortAgainst]
		for indexToSort := indexToSortAgainst + 1; indexToSort < len(list); indexToSort++ {
			if lt_comparator(list[indexToSort], valueToSwap) {
				indexToSwap = indexToSort
				valueToSwap = list[indexToSort]
			}
		}

		tmp := list[indexToSortAgainst]
		list[indexToSortAgainst] = valueToSwap
		list[indexToSwap] = tmp
	}

	return list
}

func MergeSort[N generics.Sortable](list []N) []N {
	return MergeSortWithComparator(list, generics.ComparatorLessThan[N])
}

func MergeSortWithComparator[T any](list []T, lt_comparator func(T, T) bool) []T {
	if len(list) <= 1 {
		return list
	}

	centerIndex := len(list) / 2
	leftHalf := MergeSortWithComparator(list[:centerIndex], lt_comparator)
	rightHalf := MergeSortWithComparator(list[centerIndex:], lt_comparator)

	return MergeWithComparator(leftHalf, rightHalf, lt_comparator)
}

func MergeWithComparator[T any](leftHalf, rightHalf []T, lt_comparator func(T, T) bool) []T {
	output := []T{}
	leftIndex := 0
	rightIndex := 0

	/* Mergeuntil one list exhausted */
	for leftIndex < len(leftHalf) && rightIndex < len(rightHalf) {
		if lt_comparator(leftHalf[leftIndex], rightHalf[rightIndex]) {
			output = append(output, leftHalf[leftIndex])
			leftIndex++
		} else {
			output = append(output, rightHalf[rightIndex])
			rightIndex++
		}
	}

	/* Append the remains of the unexhausted list */
	if leftIndex < len(leftHalf) {
		output = append(output, leftHalf[leftIndex:]...)
	} else if rightIndex < len(rightHalf) {
		output = append(output, rightHalf[rightIndex:]...)
	}

	return output
}

var MERGE_SORT_THREAD_CAP = 8

func MergeSortThreaded[N generics.Sortable](list []N) []N {
	return MergeSortThreadedWithComparator(list, generics.ComparatorLessThan[N])
}

func MergeSortThreadedWithComparator[T any](list []T, lt_comparator func(T, T) bool) []T {
	if len(list) <= 1 {
		return list
	}

	var helper func(list []T, ch chan []T) []T
	semaphore := make(chan int, MERGE_SORT_THREAD_CAP/2)

	helper = func(list []T, ch chan []T) []T {
		if len(list) <= 1 {
			if ch != nil {
				ch <- list
			}
			return list
		}

		centerIndex := len(list) / 2
		var leftHalf []T
		var rightHalf []T

		select {
		case semaphore <- 1:
			new_ch := make(chan []T)
			go helper(list[:centerIndex], new_ch)
			go helper(list[centerIndex:], new_ch)
			leftHalf, rightHalf = <-new_ch, <-new_ch
		default:
			leftHalf = helper(list[:centerIndex], nil)
			rightHalf = helper(list[centerIndex:], nil)
		}

		res := MergeWithComparator(leftHalf, rightHalf, lt_comparator)

		if ch != nil {
			ch <- res
		}

		return res
	}

	return helper(list, nil)
}

func QuickSort[S generics.Sortable](list []S) []S {
	return QuickSortWithComparator(list, generics.ComparatorLesserEqual[S])
}

func QuickSortWithComparator[T any](list []T, le_comparator func(T, T) bool) []T {

	var partition func(int, int) int
	partition = func(start int, end int) int {
		pivot := list[end]
		slotToSwapInto := start - 1

		for curIndex := start; curIndex < end; curIndex++ {
			if le_comparator(list[curIndex], pivot) {
				slotToSwapInto++
				tmp := list[slotToSwapInto]
				list[slotToSwapInto] = list[curIndex]
				list[curIndex] = tmp
			}
		}

		list[end] = list[slotToSwapInto+1]
		list[slotToSwapInto+1] = pivot

		return slotToSwapInto + 1
	}

	var helper func(int, int)
	helper = func(start int, end int) {
		if start < end {
			pivotIndex := partition(start, end)
			helper(start, pivotIndex-1)
			helper(pivotIndex+1, end)
		}
	}

	helper(0, len(list)-1)
	return list
}

var QUICK_SORT_THREAD_CAP int = 4

func QuickSortThreaded[S generics.Sortable](list []S) []S {
	return QuickSortThreadedWithComparator(list, generics.ComparatorLesserEqual[S])
}

func QuickSortThreadedWithComparator[T any](list []T, le_comparator func(T, T) bool) []T {

	chunkSize := len(list) / QUICK_SORT_THREAD_CAP

	if chunkSize < 4 {
		return QuickSortWithComparator(list, le_comparator)
	}

	subsets := [][]T{}
	var start int
	for i := 0; i < QUICK_SORT_THREAD_CAP; i++ {
		start = i * chunkSize
		subsets = append(subsets, list[start:start+chunkSize])
	}
	subsets[0] = append(subsets[0], list[start+chunkSize:]...)

	channels := make([]chan []T, len(subsets))

	for i, sub := range subsets {
		var sub_cpy []T
		copy(sub_cpy, sub)
		channels[i] = make(chan []T)

		go func(ch chan []T) {
			ch <- QuickSortWithComparator(sub_cpy, le_comparator)
		}(channels[i])
	}

	output := []T{}
	for _, ch := range channels {
		res := <-ch
		output = MergeWithComparator(res, output, le_comparator)
	}

	return output
}
