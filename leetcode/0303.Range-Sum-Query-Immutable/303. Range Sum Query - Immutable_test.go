package leetcode

import (
	"testing"
)

func Benchmark_Problem303(b *testing.B) {
	for bbe := 0; bbe < b.N; bbe++ {
		obj := Constructor303([]int{-2, 0, 3, -5, 2, -1})

		obj.SumRange(0, 2) // return 1
		obj.SumRange(2, 5) // return -1
		obj.SumRange(0, 5) // return -3
	}
}
