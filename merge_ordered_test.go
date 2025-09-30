package iter_lib_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lezhnev74/iter_lib"
)

func TestMergeOrdered2(t *testing.T) {
	type args struct {
		s1   iter.Seq2[string, int]
		s2   iter.Seq2[string, int]
		desc bool
	}
	type testCase struct {
		name string
		args args
		want []kv
	}
	tests := []testCase{
		{
			"empty both",
			args{
				s1:   kvsToSeq([]kv(nil)),
				s2:   kvsToSeq([]kv(nil)),
				desc: false,
			},
			[]kv(nil),
		},
		{
			"empty first",
			args{
				s1:   kvsToSeq([]kv(nil)),
				s2:   kvsToSeq([]kv{{"a", 1}, {"b", 2}, {"c", 3}}),
				desc: false,
			},
			[]kv{{"a", 1}, {"b", 2}, {"c", 3}},
		},
		{
			"empty second",
			args{
				s1:   kvsToSeq([]kv{{"a", 1}, {"b", 2}, {"c", 3}}),
				s2:   kvsToSeq([]kv(nil)),
				desc: false,
			},
			[]kv{{"a", 1}, {"b", 2}, {"c", 3}},
		},
		{
			"merge ascending",
			args{
				s1:   kvsToSeq([]kv{{"a", 1}, {"b", 2}, {"c", 3}}),
				s2:   kvsToSeq([]kv{{"d", 4}, {"e", 5}, {"f", 6}}),
				desc: false,
			},
			[]kv{{"a", 1}, {"b", 2}, {"c", 3}, {"d", 4}, {"e", 5}, {"f", 6}},
		},
		{
			"merge descending",
			args{
				s1:   kvsToSeq([]kv{{"c", 3}, {"b", 2}, {"a", 1}}),
				s2:   kvsToSeq([]kv{{"f", 6}, {"e", 5}, {"d", 4}}),
				desc: true,
			},
			[]kv{{"f", 6}, {"e", 5}, {"d", 4}, {"c", 3}, {"b", 2}, {"a", 1}},
		},
		{
			"merge descending with dups",
			args{
				s1:   kvsToSeq([]kv{{"a", 1}}),
				s2:   kvsToSeq([]kv{{"a", 1}}),
				desc: true,
			},
			[]kv{{"a", 1}, {"a", 1}},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotIt := iter_lib.MergeOrdered2(tt.args.s1, tt.args.s2, tt.args.desc)
				results := seq2ToKvs(gotIt)
				require.Equal(
					t, tt.want, results,
					"MergeOrdered2(%v, %v, %v) = %v, want %v",
					tt.args.s1, tt.args.s2, tt.args.desc, results, tt.want,
				)
			},
		)
	}
}

func TestMergeOrdered(t *testing.T) {
	type args struct {
		s1   iter.Seq[int]
		s2   iter.Seq[int]
		desc bool
	}
	type testCase struct {
		name string
		args args
		want []int
	}
	tests := []testCase{
		{
			"empty both",
			args{
				s1:   slices.Values([]int{}),
				s2:   slices.Values([]int{}),
				desc: false,
			},
			[]int(nil),
		},
		{
			"empty first",
			args{
				s1:   slices.Values([]int{}),
				s2:   slices.Values([]int{1, 2, 3}),
				desc: false,
			},
			[]int{1, 2, 3},
		},
		{
			"empty second",
			args{
				s1:   slices.Values([]int{1, 2, 3}),
				s2:   slices.Values([]int{}),
				desc: false,
			},
			[]int{1, 2, 3},
		},
		{
			"merge ascending",
			args{
				s1:   slices.Values([]int{1, 3, 5}),
				s2:   slices.Values([]int{2, 4, 6}),
				desc: false,
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"merge descending",
			args{
				s1:   slices.Values([]int{5, 3, 1}),
				s2:   slices.Values([]int{6, 4, 2}),
				desc: true,
			},
			[]int{6, 5, 4, 3, 2, 1},
		},
		{
			"merge descending sparse",
			args{
				s1:   slices.Values([]int{6, 5, 1}),
				s2:   slices.Values([]int{4, 3, 2}),
				desc: true,
			},
			[]int{6, 5, 4, 3, 2, 1},
		},
		{
			"merge descending dup",
			args{
				s1:   slices.Values([]int{2, 1}),
				s2:   slices.Values([]int{1, 0}),
				desc: true,
			},
			[]int{2, 1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				gotIt := iter_lib.MergeOrdered(tt.args.s1, tt.args.s2, tt.args.desc)
				results := slices.Collect(gotIt)
				require.Equal(
					t, tt.want, results,
					"MergeOrdered(%v, %v, %v) = %v, want %v",
					tt.args.s1, tt.args.s2, tt.args.desc, results, tt.want,
				)
			},
		)
	}
}
