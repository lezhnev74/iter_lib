package iter_lib_test

import (
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	iter_lib "github.com/lezhnev74/iter"
)

func TestFilter(t *testing.T) {
	type args struct {
		input iter.Seq[int]
		f     func(int) bool
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
				f:     func(int) bool { return true },
			},
			[]int(nil),
		},
		{
			"all match",
			args{
				input: slices.Values([]int{1, 2, 3}),
				f:     func(int) bool { return true },
			},
			[]int{1, 2, 3},
		},
		{
			"none match",
			args{
				input: slices.Values([]int{1, 2, 3}),
				f:     func(int) bool { return false },
			},
			[]int(nil),
		},
		{
			"one match",
			args{
				input: slices.Values([]int{1, 2, 3}),
				f:     func(i int) bool { return i == 1 },
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotIt := iter_lib.Filter(tt.args.input, tt.args.f)
				results := slices.Collect(gotIt)
				require.Equal(
					t, tt.want, results,
					"Filter(%v, %v) = %v, want %v",
					tt.args.input, tt.args.f, results, tt.want,
				)
			},
		)
	}
}

func TestFilter2(t *testing.T) {
	type args struct {
		input iter.Seq2[int, int]
		f     func(int, int) bool
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
				f:     func(int, int) bool { return true },
			},
			map[int]int{},
		},
		{
			"all match",
			args{
				input: slices.All([]int{1, 2}),
				f:     func(int, int) bool { return true },
			},
			map[int]int{0: 1, 1: 2},
		},
		{
			"none match",
			args{
				input: slices.All([]int{1, 2}),
				f:     func(int, int) bool { return false },
			},
			map[int]int{},
		},
		{
			"some match",
			args{
				input: slices.All([]int{1, 2, 3}),
				f:     func(k, v int) bool { return k == 1 || v == 1 },
			},
			map[int]int{0: 1, 1: 2},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotIt := iter_lib.Filter2(tt.args.input, tt.args.f)
				results := maps.Collect(gotIt)
				require.Equal(
					t, tt.want, results,
					"Filter2(%v, %v) = %v, want %v",
					tt.args.input, tt.args.f, results, tt.want,
				)
			},
		)
	}
}
