package generics

func ComparatorLessThan[N Sortable](v1 N, v2 N) bool {
	return v1 < v2
}

func ComparatorGreaterThan[N Sortable](v1 N, v2 N) bool {
	return v1 > v2
}

func ComparatorGreaterEqual[N Sortable](v1 N, v2 N) bool {
	return v1 >= v2
}

func ComparatorLesserEqual[N Sortable](v1 N, v2 N) bool {
	return v1 <= v2
}

func ComparatorEquals[N comparable](v1 N, v2 N) bool {
	return v1 == v2
}

func ComparatorNotEquals[N comparable](v1 N, v2 N) bool {
	return v1 != v2
}
