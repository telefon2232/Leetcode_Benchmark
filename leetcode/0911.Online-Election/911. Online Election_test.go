package leetcode

import (
	"testing"
)

func Benchmark_Problem911(b *testing.B) {
	obj := Constructor911([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	(obj.Q(3))  // 0
	(obj.Q(12)) // 1
	(obj.Q(25)) // 1
	(obj.Q(15)) // 0
	(obj.Q(24)) // 0
	(obj.Q(8))  // 1

	obj = Constructor911([]int{0, 0, 0, 0, 1}, []int{0, 6, 39, 52, 75})
	(obj.Q(45)) // 0
	(obj.Q(49)) // 0
	(obj.Q(59)) // 0
	(obj.Q(68)) // 0
	(obj.Q(42)) // 0
	(obj.Q(37)) // 0
	(obj.Q(99)) // 0
	(obj.Q(26)) // 0
	(obj.Q(78)) // 0
	(obj.Q(43)) // 0
}
