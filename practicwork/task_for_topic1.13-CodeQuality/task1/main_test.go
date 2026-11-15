package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		b       float64
		want    float64
		wantErr bool
	}{
		{name: "Деление на 0", a: 5.4, b: 0, want: 0, wantErr: true},
		{name: "Обычное деление", a: 8, b: 0.5, want: 16, wantErr: false},
		{name: "Деление с отрицательным знаком", a: 54, b: -32, want: -1.6875, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Divide(tt.a, tt.b)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Divide() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Divide() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(8, 0.5)
	}
}
