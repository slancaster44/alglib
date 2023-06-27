package arrutil

import (
	"testing"
)

func cmpArr[C comparable, A []C](val1, val2 A) bool {
	if len(val1) != len(val2) {
		return false
	}

	for index, val1_v := range val1 {
		if val2[index] != val1_v {
			return false
		}
	}

	return true
}

func TestMaxSubArr(t *testing.T) {
	test1 := []int{1, -3, 4, 5, -33, 11}
	test2 := []float64{1.1, -33, 11, 44}
	test3 := []int{1, 2, 3, 4}
	test4 := []int{}
	test5 := []int{11}

	if !cmpArr(MaxiumSubArray(test1), []int{11}) {
		t.Log("Input:", test1, "Output:", MaxiumSubArray(test1))
		t.Error("[FAIL] Test 1")
	} else if !cmpArr(MaxiumSubArray(test2), []float64{11, 44}) {
		t.Log("Input:", test2, "Output:", MaxiumSubArray(test2))
		t.Error("[FAIL] Test 2")
	} else if !cmpArr(MaxiumSubArray(test3), []int{1, 2, 3, 4}) {
		t.Log("Input:", test3, "Output:", MaxiumSubArray(test3))
		t.Error("[FAIL] Test 3")
	} else if !cmpArr(MaxiumSubArray(test4), []int{}) {
		t.Log("Input:", test4, "Output:", MaxiumSubArray(test4))
		t.Error("[FAIL] Test 4")
	} else if !cmpArr(MaxiumSubArray(test5), []int{11}) {
		t.Log("Input:", test5, "Output:", MaxiumSubArray(test5))
		t.Error("[FAIL] Test 5")
	}

}
