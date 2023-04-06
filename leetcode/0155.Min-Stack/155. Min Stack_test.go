package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem155(b *testing.B) {
	obj1 := Constructor155()
	obj1.Push(1)

	obj1.Push(0)

	obj1.Push(10)

	obj1.Pop()

	param3 := obj1.Top()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj1.GetMin()
	fmt.Printf("param_4 = %v\n", param4)
}
