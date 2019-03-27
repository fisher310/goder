package basic

import "testing"

func TestCalcTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := CalcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("CalcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkCalcTriangle(b *testing.B) {
	a1, b1, c1 := 30000, 40000, 50000
	for i := 0; i < b.N; i++ {
		if c1 != CalcTriangle(a1, b1) {
			b.Errorf("got error")
		}
	}
}
