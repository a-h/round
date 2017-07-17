package round

import "math"

// ToEven rounds to the nearest even number.
// Used in Python 3 (but not Python 2, which uses AwayFromZero).
// -3.5 rounds to -4
// -2.5 rounds to -2
// -1.5 rounds to -2
// 1.5 rounds to 2
// 2.5 rounds to 2
// 3.5 rounds to 4
func ToEven(v float64, decimals int) float64 {
	if math.IsNaN(v) {
		return math.NaN()
	}

	// Round by multiplying by 10 then flooring the result.
	// e.g. 1.234 to 2 dp would multiply by 100 to get 123.4.
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}

	// This conversion floors the float, e.g. 123.4 returns 123.
	// We can undo the operation later to carry out the rounding.
	vpow := v * pow
	intPart, fracPart := math.Modf(vpow)

	// If we're at the midpoint.
	m := midpointPositive
	if v < 0 {
		m = midpointNegative
	}

	if isWithin(fracPart, m, tolerance) {
		if int64(intPart)%2 == 0 {
			// It's even, so round to it.
			return intPart / pow
		}
	}

	// Or we need to round to the nearest even, for negative
	// numbers, that's by subtraction, for positive, by
	// addition.
	return float64(int64(vpow+(m*1.001))) / pow
}

func isWithin(a float64, b float64, tolerance float64) bool {
	if a > b {
		return a-b <= tolerance
	}
	return b-a <= tolerance
}
