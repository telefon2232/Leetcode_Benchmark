package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem338(b *testing.B) {
	obj := Constructor([]*NestedInteger{})
	fmt.Printf("obj = %v\n", obj)
	fmt.Printf("obj = %v\n", obj.Next())
	fmt.Printf("obj = %v\n", obj.HasNext())
}
