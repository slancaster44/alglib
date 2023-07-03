package utility

func Arrcmp[C comparable, A []C](val1, val2 A) bool {
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
