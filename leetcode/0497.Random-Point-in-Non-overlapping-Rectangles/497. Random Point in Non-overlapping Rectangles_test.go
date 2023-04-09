package leetcode

import (
	"testing"
)

func Benchmark_Problem497(b *testing.B) {
	w := [][]int{{1, 1, 5, 5}}
	sol := Constructor497(w)
	for bbe := 0; bbe < b.N; bbe++ {
		sol.Pick()
		sol.Pick()
		sol.Pick()
		sol.Pick()
		sol.Pick()
		sol.Pick()
	}
}
