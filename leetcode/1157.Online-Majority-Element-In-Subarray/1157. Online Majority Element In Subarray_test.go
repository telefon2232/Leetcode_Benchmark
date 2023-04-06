package leetcode

import (
	"testing"
)

func Benchmark_Problem1157(b *testing.B) {
	arr := []int{1, 1, 2, 2, 1, 1}
	obj := Constructor1157(arr)

	obj.Query(0, 5, 4) //1
	obj.Query(0, 3, 3) //-1
	obj.Query(2, 3, 2) //2
}
