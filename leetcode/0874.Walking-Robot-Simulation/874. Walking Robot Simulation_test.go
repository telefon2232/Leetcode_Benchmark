package leetcode

import "testing"

func Benchmark_robotSim(b *testing.B) {
	type args struct {
		commands  []int
		obstacles [][]int
	}
	cases := []struct {
		name string
		args args
		want int
	}{

		{
			"case 2",
			args{
				commands:  []int{4, -1, 4, -2, 4},
				obstacles: [][]int{{2, 4}},
			},
			65,
		},
	}
	for bbe := 0; bbe < b.N; bbe++ {
		for _, tt := range cases {
			b.Run(tt.name, func(b *testing.B) {
				if got := robotSim(tt.args.commands, tt.args.obstacles); got != tt.want {

				}
			})
		}
	}
}
