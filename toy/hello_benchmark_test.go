package toy

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}
func mu(a, b int) int {
	return a * b
}
func addmore(a, b, c, d, e, f, g int) int {
	return a + b + c + d + e + f + g
}
func Benchmark_A(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}
func Benchmark_addmore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addmore(1, 2, 3, 4, 5, 6, 7)
	}
}
func Benchmark_M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mu(1, 2)
	}
}
