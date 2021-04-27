package fibonachi

import "testing"

func TestFibVars(t *testing.T) {
	insideTestFib(t, FibVars, "FibVars")
}

func TestFibCache(t *testing.T) {
	insideTestFib(t, FibCache, "FibCache")
}

func TestFibRecur(t *testing.T) {
	insideTestFib(t, FibRecur, "FibRecur")
}

func insideTestFib(t *testing.T, f func(int) int, fn string) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantF   int
		wantErr bool
	}{
		{"-2", args{-2}, -2, true},
		{"0", args{0}, 0, false},
		{"1", args{1}, 1, false},
		{"2", args{2}, 1, false},
		{"53", args{30}, 832040, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF := f(tt.args.n)
			if gotF != tt.wantF {
				t.Errorf("%s() = %v, want %v", fn, gotF, tt.wantF)
			}
		})
	}
}

var globalF int

func BenchmarkFibVars(t *testing.B) {
	var f int
	for i := 0; i <= t.N; i++ {
		f = FibVars(30)
	}

	globalF = f
}
func BenchmarkFibCache(t *testing.B) {
	var f int
	for i := 0; i <= t.N; i++ {
		f = FibCache(30)
	}

	globalF = f
}

func BenchmarkFibRecur(t *testing.B) {
	var f int
	for i := 0; i <= t.N; i++ {
		f = FibRecur(30)
	}

	globalF = f
}
