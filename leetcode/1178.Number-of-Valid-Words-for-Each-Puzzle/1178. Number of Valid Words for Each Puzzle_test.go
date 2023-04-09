package leetcode

import (
	"reflect"
	"testing"
)

func Benchmark_findNumOfValidWords(b *testing.B) {

	words1 := []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}
	puzzles1 := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}

	type args struct {
		words   []string
		puzzles []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{words: words1, puzzles: puzzles1}, []int{1, 1, 3, 2, 4, 0}},
	}
	for bbe := 0; bbe < b.N; bbe++ {
		for _, tt := range tests {
			b.Run(tt.name, func(b *testing.B) {
				if got := findNumOfValidWords(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
					b.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
				}
			})
		}
	}

}
