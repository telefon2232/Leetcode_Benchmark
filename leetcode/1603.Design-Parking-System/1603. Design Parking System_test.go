package leetcode

import (
	"testing"
)

func Benchmark_Problem1603(b *testing.B) {
	obj := Constructor(1, 1, 0)
	for bbe := 0; bbe < b.N; bbe++ {
		obj.AddCar(1)
		obj.AddCar(2)
		obj.AddCar(3)
		obj.AddCar(1)
	}
}
