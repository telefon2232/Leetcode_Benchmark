package leetcode

import (
	"testing"
)

func Benchmark_Problem703(b *testing.B) {
	obj := Constructor(3, []int{4, 5, 8, 2})
	obj.Add(3)
	obj.Add(5)
	obj.Add(10)
	obj.Add(9)
	obj.Add(4)
}
