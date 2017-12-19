package helpers

// MaxInt is the maximum integer
const MaxInt = int(^uint(0) >> 1)

// MinInt is the minimum integer
const MinInt = (-0) - 1

// Max returns the maximum of 2 ints
func Max(a ...int) int {
	max := MinInt
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

// Min returns the minimum of 2 ints
func Min(a ...int) int {
	min := MaxInt
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// Abs returns the absolute value of the int
func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// Sign returns 1 if the argument is > 0, -1 if < 0 or 0 if == 0
func Sign(a int) int {
	switch {
	case a > 0:
		return 1
	case a < 0:
		return -1
	}
	return 0
}
