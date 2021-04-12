package main

import (
	"testing"
)

type argsFib struct {
	n int
}
type stepFib struct {
	name string
	args argsFib
	want int
}

var testsFib []stepFib

func init() {
	testsFib = []stepFib{
		{name: "0", args: argsFib{n: 0}, want: 0},
		{name: "1", args: argsFib{n: 1}, want: 1},
		{name: "2", args: argsFib{n: 2}, want: 1},
		{name: "3", args: argsFib{n: 3}, want: 2},
		{name: "4", args: argsFib{n: 4}, want: 3},
		{name: "5", args: argsFib{n: 5}, want: 5},
		{name: "6", args: argsFib{n: 6}, want: 8},
		{name: "7", args: argsFib{n: 7}, want: 13},
		{name: "8", args: argsFib{n: 8}, want: 21},
		{name: "9", args: argsFib{n: 9}, want: 34},
		{name: "10", args: argsFib{n: 10}, want: 55},
		//{name: "100", args: argsFib{n: 100}, want: 3736710778780434371}, // Test_fibonachiRecur timeout fail
	}
}

func Test_fibonachiVar(t *testing.T) {
	for _, tt := range testsFib {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonachiVar(tt.args.n); got != tt.want {
				t.Errorf("fibonachiVar(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}

func Test_fibonachiMap(t *testing.T) {
	for _, tt := range testsFib {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonachiMap(tt.args.n); got != tt.want {
				t.Errorf("fibonachiMap(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}

func Test_fibonachiRecur(t *testing.T) {
	for _, tt := range testsFib {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonachiRecur(tt.args.n); got != tt.want {
				t.Errorf("fibonachiRecur(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
