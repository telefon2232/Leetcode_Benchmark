package leetcode

import (
	"fmt"
	"testing"
	"unsafe"
)

type question419 struct {
	para419
	ans419
}

// para 是参数
// one 代表第一个参数
type para419 struct {
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans419 struct {
	one int
}

func Benchmark_Problem419(b *testing.B) {

	qs := []question419{

		{
			para419{[][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}},
			ans419{2},
		},

		{
			para419{[][]byte{{'.'}}},
			ans419{0},
		},
	}

	for bbe := 0; bbe < b.N; bbe++ {
		for _, q := range qs {
			{
				_, p := q.ans419, q.para419
				bytesArrayToStringArray(p.one)
				countBattleships(p.one)
			}

		}
	}
}

// 在运行go test时 为了更直观地显示[][]byte中的字符而非ASCII码数值
// bytesArrayToStringArray converts [][]byte to []string
func bytesArrayToStringArray(b [][]byte) []string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = fmt.Sprintf("[%v]", *(*string)(unsafe.Pointer(&b[i])))
	}
	return s
}
