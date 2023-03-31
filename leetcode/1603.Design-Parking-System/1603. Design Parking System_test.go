package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem1603(b *testing.B) {
	obj := Constructor(1, 1, 0)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("obj = %v\n", obj.AddCar(1))
	fmt.Printf("obj = %v\n", obj.AddCar(2))
	fmt.Printf("obj = %v\n", obj.AddCar(3))
	fmt.Printf("obj = %v\n", obj.AddCar(1))
}
