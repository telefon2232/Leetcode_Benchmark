package leetcode

import (
	"fmt"
	"testing"
)

func Benchmark_Problem732(b *testing.B) {
	obj := Constructor732()
	fmt.Printf("book = %v\n\n", obj.Book(10, 20)) // returns 1
	fmt.Printf("book = %v\n\n", obj.Book(50, 60)) // returns 1
	fmt.Printf("book = %v\n\n", obj.Book(10, 40)) // returns 2
	fmt.Printf("book = %v\n\n", obj.Book(5, 15))  // returns 3
	fmt.Printf("book = %v\n\n", obj.Book(5, 10))  // returns 3
	fmt.Printf("book = %v\n\n", obj.Book(25, 55)) // returns 3
}
