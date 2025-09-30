package iter_lib_test

import (
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lezhnev74/iter_lib"
)

func TestMap(t *testing.T) {
	type args struct {
		input iter.Seq[int]
		f     func(int) int
	}
	type testCase struct {
		name string
		args args
		want []int
	}
	tests := []testCase{
		{
			"empty",
			args{
				input: slices.Values([]int{}),
				f:     func(i int) int { return i },
			},
			[]int(nil),
		},
		{
			"identity",
			args{
				input: slices.Values([]int{1, 2, 3}),
				f:     func(i int) int { return i },
			},
			[]int{1, 2, 3},
		},
		{
			"double",
			args{
				input: slices.Values([]int{1, 2, 3}),
				f:     func(i int) int { return i * 2 },
			},
			[]int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotIt := iter_lib.Map(tt.args.input, tt.args.f)
				results := slices.Collect(gotIt)
				require.Equal(
					t, tt.want, results,
					"Map(%v, %v) = %v, want %v",
					tt.args.input, tt.args.f, results, tt.want,
				)
			},
		)
	}
}

func TestMap2(t *testing.T) {
	type args struct {
		input iter.Seq2[int, int]
		f     func(int) int
	}
	type testCase struct {
		name string
		args args
		want map[int]int
	}
	tests := []testCase{
		{
			"empty",
			args{
				input: slices.All([]int{}),
				f:     func(i int) int { return i },
			},
			map[int]int{},
		},
		{
			"identity",
			args{
				input: slices.All([]int{1, 2}),
				f:     func(i int) int { return i },
			},
			map[int]int{0: 1, 1: 2},
		},
		{
			"double",
			args{
				input: slices.All([]int{1, 2}),
				f:     func(i int) int { return i * 2 },
			},
			map[int]int{0: 2, 1: 4},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotIt := iter_lib.Map2(tt.args.input, tt.args.f)
				results := maps.Collect(gotIt)
				require.Equal(
					t, tt.want, results,
					"Map2(%v, %v) = %v, want %v",
					tt.args.input, tt.args.f, results, tt.want,
				)
			},
		)
	}
}
