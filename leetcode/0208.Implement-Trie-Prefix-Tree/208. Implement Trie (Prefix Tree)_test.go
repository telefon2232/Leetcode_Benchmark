package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem208(b *testing.B) {
	obj := Constructor208()

	obj.Insert("apple")

	param1 := obj.Search("apple")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param2 := obj.Search("app")
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
	param3 := obj.StartsWith("app")
	fmt.Printf("param_3 = %v obj = %v\n", param3, obj)
	obj.Insert("app")

	param4 := obj.Search("app")
	fmt.Printf("param_4 = %v obj = %v\n", param4, obj)
}
