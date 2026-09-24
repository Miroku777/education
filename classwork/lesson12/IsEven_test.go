//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out
//go tool cover -html=coverage.out -o coverage.html

package main

import (
	"testing"
)

func BenchmarkInsertInSliceZero100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInSliceZero(100)
	}
}
func BenchmarkInsertInSliceZero10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInSliceZero(10000)
	}
}
func BenchmarkInsertInSliceZero100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInSliceZero(100000)
	}
}
func BenchmarkInsertInSliceV100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInSliceV(100)
	}
}
func BenchmarkInsertInSliceV10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInSliceV(10000)
	}
}
func BenchmarkInsertInSliceV100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertInSliceV(100000)
	}
}

func TestIsEven(t *testing.T) {
	var testCases = []struct {
		description string
		input       int
		want        string
	}{
		{"input -1", -1, "No"},
		{"input 2", 2, "Yes"},
		{"input 0", 0, "Yes"},
		{"input -100", -100, "Yes"},
	}
	for _, tst := range testCases {
		t.Run(tst.description, func(t *testing.T) {
			res := IsEven(tst.input)
			if res != tst.want {
				t.Errorf("Got: %s, want: %s", res, tst.want)
			}
		})
	}
}
