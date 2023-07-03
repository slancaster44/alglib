package alglib

import (
	"alglib/utility"
	"math/rand"
	"testing"
	"time"
)

const TEST_ARR_LEN int = 8192 * 2

func TestIsSorted(t *testing.T) {
	if !IsSorted([]int{1, 2, 3}) {
		t.Error("IsSorted() Test Failure")
	}
}

func getRandList(size int, limit int) []int {
	out := []int{}
	for i := 0; i < size; i++ {
		out = append(out, rand.Int()%limit)
	}

	return out
}

func TestIntegerSorts(t *testing.T) {
	rand.Seed(time.Now().Unix())

	arr := getRandList(TEST_ARR_LEN, 100)
	if !IsSorted(MergeSort(arr)) {
		t.Error("Merge Integer Sort Fail. Input:", arr, "Output:", MergeSort(arr))
	}

	arr = getRandList(TEST_ARR_LEN, 100)
	if !IsSorted(QuickSort(arr)) {
		t.Error("Quick integer sort fail. Input:", arr, "Output:", QuickSort(arr))
	}

	arr = getRandList(TEST_ARR_LEN, 100)
	if !IsSorted(MergeSortThreaded(arr)) {
		t.Error("Merge Integer Sort Fail. Input:", arr, "Output:", MergeSort(arr))
	}

	arr = getRandList(TEST_ARR_LEN, 100)
	res := QuickSortThreaded(arr)
	if !IsSorted(res) {
		t.Error("Threaded QuickSort Fail. Input:", arr, "Output:", res)
	}

	arr = getRandList(TEST_ARR_LEN, 100)
	if !IsSorted(InsertionSort(arr)) {
		t.Error("Insertion Integer Sort Fail. Input:", arr, "Output:", InsertionSort(arr))
	}
}

func TestPreSorted(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	if !IsSorted(MergeSort(arr)) {
		t.Error("Pre sorted merge sort fail. Input:", arr, "Output:", MergeSort(arr))
	}

	if !IsSorted(InsertionSort(arr)) {
		t.Error("Pre sorted insertion sort fail. Input:", arr, "Output:", InsertionSort(arr))
	}

	if !IsSorted(MergeSortThreaded(arr)) {
		t.Error("Merge Pre Sorted Fail. Input:", arr, "Output:", MergeSort(arr))
	}

	if !IsSorted(QuickSortThreaded(arr)) {
		t.Error("Merge Pre Sorted Fail. Input:", arr, "Output:", QuickSortThreaded(arr))
	}

}

func TestEmptyArr(t *testing.T) {
	if !utility.Arrcmp(MergeSort([]int{}), []int{}) {
		t.Error("Empty merge")
	}

	if !utility.Arrcmp(InsertionSort([]int{}), []int{}) {
		t.Error("Empty insertion")
	}

	if !IsSorted(MergeSortThreaded([]int{})) {
		t.Error("Merge Sort Empty")
	}

	if !IsSorted(QuickSortThreaded([]int{})) {
		t.Error("Quicksort threaded failed")
	}
}

var benchListSize int = 8 << 16

func getBenchmarkList() []float64 {
	output := []float64{}
	for i := 0; i < benchListSize; i++ {
		output = append(output, rand.ExpFloat64())
	}

	return output
}
func BenchmarkMergeSort(b *testing.B) {
	MergeSort(getBenchmarkList())
}

func BenchmarkMergeSortThreaded(b *testing.B) {
	MergeSortThreaded(getBenchmarkList())
}

func BenchmarkInsertionSort(b *testing.B) {
	InsertionSort(getBenchmarkList())
}

func BenchmarkQuicksort(b *testing.B) {
	QuickSort(getBenchmarkList())
}

func BenchmarkQuickSortThreaded(b *testing.B) {
	QuickSortThreaded(getBenchmarkList())
}
