package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

type question191 struct {
	para191
	ans191
}

// para 是参数
// one 代表第一个参数
type para191 struct {
	one uint32
}

// ans 是答案
// one 代表第一个答案
type ans191 struct {
	one int
}

func Benchmark_Problem191(b *testing.B) {

	qs := []question191{

		{
			para191{5},
			ans191{1},
		},
		{
			para191{13},
			ans191{2},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans191, q.para191
				input := strconv.FormatUint(uint64(p.one), 2) // 32位无符号整数转换为二进制字符串
				input = fmt.Sprintf("%0*v", 32, input)        // 格式化输出32位,保留前置0
				fmt.Printf("【input】:%v       【output】:%v\n", input, hammingWeight(p.one))
			}
		}
	}
}
