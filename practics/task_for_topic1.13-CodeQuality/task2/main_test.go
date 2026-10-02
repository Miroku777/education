package main

import "testing"

func TestSumAll(t *testing.T) {
	tests := []struct {
		name string
		line []int
		want int
	}{
		{name: "Первый тест", line: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 45},
		{name: "Второй тест", line: []int{1, 2, 3}, want: 6},
		{name: "Третий тест", line: []int{10, -2, 4, 7}, want: 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumAll(tt.line...)
			if got != tt.want {
				t.Errorf("sumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}
