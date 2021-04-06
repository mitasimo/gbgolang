package main

import "testing"

func Test_quickSort(t *testing.T) {
	tests := []struct {
		name          string
		input, sorted []int
	}{
		{
			name:   "simple",
			input:  []int{2, 4, 7, 9, 0, 5, 1},
			sorted: []int{0, 1, 2, 4, 5, 7, 9},
		},
		{
			name:   "one",
			input:  []int{2},
			sorted: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.input)
			for i, v := range tt.input {
				if v != tt.sorted[i] {
					t.Fatalf("Array %v after sorting must be %v", tt.input, tt.sorted)
				}
			}
		})
	}
}
