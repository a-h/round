package round

import "testing"

func benchmarkToEven(v float64, places int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		ToEven(v, places)
	}
}

func BenchmarkToEvenPositiveOdd(b *testing.B)  { benchmarkToEven(1.5, 0, b) }
func BenchmarkToEvenPositiveEven(b *testing.B) { benchmarkToEven(2.5, 0, b) }
