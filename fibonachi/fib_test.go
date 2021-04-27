package fibonachi

import "testing"

func TestFibVars(t *testing.T) {
	insideTestFib(t, FibVars, "FibVars")
}

func TestFibSlice(t *testing.T) {
	insideTestFib(t, FibSlice, "FibCache")
}

func TestFibMap(t *testing.T) {
	insideTestFib(t, FibMap, "FibMap")
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
		{"40", args{40}, 102334155, false},
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
		f = FibVars(5000)
	}

	globalF = f
}
func BenchmarkFibSlice(t *testing.B) {
	var f int
	for i := 0; i <= t.N; i++ {
		f = FibSlice(5000)
	}

	globalF = f
}

func BenchmarkFibMap(t *testing.B) {
	var f int
	for i := 0; i <= t.N; i++ {
		f = FibMap(5000)
	}

	globalF = f
}

func BenchmarkFibRecur(t *testing.B) {
	var f int
	for i := 0; i <= t.N; i++ {
		f = FibRecur(40)
	}

	globalF = f
}
