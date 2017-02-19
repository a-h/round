package round

// ToEven rounds to the nearest even number.
// Used in Python 3 (but not Python 2, which uses AwayFromZero).
// -3.5 rounds to -4
// -2.5 rounds to -2
// -1.5 rounds to -2
// 1.5 rounds to 2
// 2.5 rounds to 2
// 3.5 rounds to 4
func ToEven(v float64, decimals int) float64 {
	// Round by multiplying by 10 then flooring the result.
	// e.g. 1.234 to 2 dp would multiply by 100 to get 123.4.
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	// This conversion floors the float, e.g. 123.4 returns 123.
	// We can undo the operation later to carry out the rounding.
	floored := int(v * pow)

	// First, get the number after the decimal place, e.g.:
	// 123.4 would return 0.4
	afterDecimal := (v * pow) - float64(floored)

	// If we're at the midpoint.
	if afterDecimal == 0.5 || afterDecimal == -0.5 {
		if floored%2 == 0 {
			// It's even, so round to it.
			return float64(floored) / pow
		}
	}

	// Or we need to round to the nearest even, for negative
	// numbers, that's by subtraction, for positive, by
	// addition.
	if v < 0 {
		return float64(int64((v*pow)-0.5)) / pow
	}
	return float64(int64((v*pow)+0.5)) / pow
}
