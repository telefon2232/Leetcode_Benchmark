package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem232(b *testing.B) {
	obj := Constructor232()
	
	obj.Push(2)
	
	obj.Push(10)
	
	param2 := obj.Pop()
	fmt.Printf("param_2 = %v\n", param2)
	param3 := obj.Peek()
	fmt.Printf("param_3 = %v\n", param3)
	param4 := obj.Empty()
	fmt.Printf("param_4 = %v\n", param4)
}
