package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem528(b *testing.B) {
	w := []int{1, 3}
	sol := Constructor528(w)
	fmt.Printf("1.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("2.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("3.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("4.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("5.PickIndex = %v\n", sol.PickIndex())
	fmt.Printf("6.PickIndex = %v\n", sol.PickIndex())
}
