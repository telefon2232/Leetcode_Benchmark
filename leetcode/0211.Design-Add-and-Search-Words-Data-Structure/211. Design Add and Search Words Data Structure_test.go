package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem211(b *testing.B) {
	obj := Constructor211()

	obj.AddWord("bad")

	obj.AddWord("dad")

	obj.AddWord("mad")

	param1 := obj.Search("pad")
	fmt.Printf("param_1 = %v obj = %v\n", param1, obj)
	param2 := obj.Search("bad")
	fmt.Printf("param_2 = %v obj = %v\n", param2, obj)
	param3 := obj.Search(".ad")
	fmt.Printf("param_3 = %v obj = %v\n", param3, obj)
	param4 := obj.Search("b..")
	fmt.Printf("param_4 = %v obj = %v\n", param4, obj)
}
