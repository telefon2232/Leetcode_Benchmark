package leetcode

import (
	"testing"
)

func Benchmark_Problem478(b *testing.B) {
	for bbe := 0; bbe < b.N; bbe++ {
		obj := Constructor(1, 0, 0)
		obj.RandPoint()
		obj.RandPoint()
		obj.RandPoint()

		obj = Constructor(10, 5, -7.5)
		obj.RandPoint()
		obj.RandPoint()
		obj.RandPoint()
	}
}
