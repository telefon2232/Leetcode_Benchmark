package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem677(b *testing.B) {
	obj := Constructor()
	fmt.Printf("obj = %v\n", obj)
	obj.Insert("apple", 3)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("obj.sum = %v\n", obj.Sum("ap"))
	obj.Insert("app", 2)
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("obj.sum = %v\n", obj.Sum("ap"))
}
