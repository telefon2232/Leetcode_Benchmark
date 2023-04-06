package leetcode

import (
	"testing"
)

func Benchmark_Problem528(b *testing.B) {
	w := []int{1, 3}
	sol := Constructor528(w)
	sol.PickIndex()
	sol.PickIndex()
	sol.PickIndex()
	sol.PickIndex()
	sol.PickIndex()
	sol.PickIndex()
}
