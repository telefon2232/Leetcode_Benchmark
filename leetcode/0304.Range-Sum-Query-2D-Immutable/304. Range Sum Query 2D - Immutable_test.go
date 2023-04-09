package leetcode

import (
	"testing"
)

func Benchmark_Problem304(b *testing.B) {
	obj := Constructor(
		[][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		},
	)
	for bbe := 0; bbe < b.N; bbe++ {
		obj.SumRegion(2, 1, 4, 3)
		obj.SumRegion(1, 1, 2, 2)
		obj.SumRegion(1, 2, 2, 4)
	}
}
