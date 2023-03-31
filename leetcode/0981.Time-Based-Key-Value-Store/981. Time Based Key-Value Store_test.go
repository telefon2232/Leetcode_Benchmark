package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem981(b *testing.B) {
	obj := Constructor981()
	obj.Set("foo", "bar", 1)
	fmt.Printf("Get = %v\n", obj.Get("foo", 1))
	fmt.Printf("Get = %v\n", obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Printf("Get = %v\n", obj.Get("foo", 4))
	fmt.Printf("Get = %v\n", obj.Get("foo", 5))
}
