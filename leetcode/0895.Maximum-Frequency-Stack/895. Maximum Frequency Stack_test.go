package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem895(b *testing.B) {
	obj := Constructor895()
	for bbe := 0; bbe < b.N; bbe++ {
		obj.Push(5)

		obj.Push(7)

		obj.Push(5)

		obj.Push(7)

		obj.Push(4)

		obj.Push(5)

		param1 := obj.Pop()
		fmt.Print(param1)
		param1 = obj.Pop()

		param1 = obj.Pop()

		param1 = obj.Pop()

		param1 = obj.Pop()

		param1 = obj.Pop()

	}
}
