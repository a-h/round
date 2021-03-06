package round

import "math"

// AwayFromZero rounds numbers away from zero.
// Used in Python 2 (but not Python 3, which uses ToEven).
// -3.5 rounds to -4
// -2.5 rounds to -3
// -1.5 rounds to -2
// 1.5 rounds to 2
// 2.5 rounds to 3
// 3.5 rounds to 4
func AwayFromZero(v float64, decimals int) float64 {
	if math.IsNaN(v) {
		return math.NaN()
	}

	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	if v < 0 {
		return float64(int((v*pow)-0.5)) / pow
	}
	return float64(int((v*pow)+0.5)) / pow
}
