package alglib

import "alglib/generics"

func BinarySearch[N generics.Sortable](sorted_list []N, query N) int {
	return BinarySearchWithComparator(sorted_list, query, generics.ComparatorLessThan[N], generics.ComparatorGreaterThan[N])
}

func BinarySearchWithComparator[T any](sorted_list []T, query T, less_than_comparator, greater_than_comparator func(T, T) bool) int {

	if len(sorted_list) == 0 {
		return -1
	}

	var helper func(lower, upper int) int

	helper = func(lower, upper int) int {
		if lower == upper {
			return -1
		}

		middle := ((upper - lower) / 2) + lower

		valueToCheck := sorted_list[middle]

		if less_than_comparator(valueToCheck, query) {
			return helper(lower, middle)
		} else if greater_than_comparator(valueToCheck, query) {
			return helper(middle, upper)
		}

		return middle
	}

	return helper(0, len(sorted_list)-1)
}

func LinearSearch[T comparable](list []T, query T) int {
	return LinearSearchWithComparator(list, query, generics.ComparatorEquals[T])
}

func LinearSearchWithComparator[T any](list []T, query T, eq_comparator func(T, T) bool) int {
	for idx, val := range list {
		if eq_comparator(val, query) {
			return idx
		}
	}

	return -1
}
