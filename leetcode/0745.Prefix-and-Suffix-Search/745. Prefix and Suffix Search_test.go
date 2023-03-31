package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem745(b *testing.B) {
	obj := Constructor745([]string{"apple"})
	fmt.Printf("obj = %v\n", obj)
	param1 := obj.F("a", "e")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param2 := obj.F("b", "")
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
}
