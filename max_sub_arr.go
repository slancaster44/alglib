package alglib

import (
	"alglib/generics"
)

func MaxiumSubArray[K generics.SignedNumber](input []K) []K {
	if len(input) == 0 {
		return input
	}

	startIndex, endIndex, _ := maxSubArrHelper(input)
	return input[startIndex : endIndex+1]
}

func maxSubArrHelper[K generics.SignedNumber](input []K) (int, int, K) {
	if len(input) == 1 {
		return 0, 0, input[0]
	}

	middle := len(input) / 2
	onDivStart, onDivEnd, onDivSum := maxAcrossDivide(input, middle)
	leftStart, leftEnd, leftSum := maxSubArrHelper(input[:middle])
	rightStart, rightEnd, rightSum := maxSubArrHelper(input[middle:])

	if onDivSum > leftSum && onDivSum > rightSum {
		return onDivStart, onDivEnd, onDivSum
	} else if leftSum > onDivSum && leftSum > rightSum {
		return middle + leftStart, middle + leftEnd, leftSum
	} else {
		return middle + rightStart, middle + rightEnd, rightSum
	}

}

func maxAcrossDivide[K generics.SignedNumber](input []K, middle int) (int, int, K) {

	leftVolume := input[middle-1]
	leftSum := input[middle-1]
	left_index := middle - 1

	for i := middle - 2; i >= 0; i-- {
		leftSum += input[i]

		if leftSum > leftVolume {
			leftVolume = leftSum
			left_index = i
		}
	}

	rightVolume := input[middle]
	rightSum := input[middle]
	right_index := middle

	for i := middle + 1; i < len(input); i++ {
		rightSum += input[i]

		if rightSum > rightVolume {
			rightVolume = rightSum
			right_index = i
		}
	}

	return left_index, right_index, leftVolume + rightVolume
}
