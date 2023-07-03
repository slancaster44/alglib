package generics

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type Float interface {
	~float32 | ~float64
}

type SignedNumber interface {
	SignedInteger | Float
}

type Number interface {
	SignedNumber | UnsignedInteger
}

type Sortable interface {
	Number | ~string
}
