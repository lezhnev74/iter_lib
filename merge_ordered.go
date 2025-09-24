package iter_lib

import (
	"cmp"
	"iter"
)

// MergeOrdered returns an iterator that yields values from both input sequences in ascending or descending order based on desc flag.
func MergeOrdered[V cmp.Ordered](s1, s2 iter.Seq[V], desc bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		next1, stop1 := iter.Pull(s1)
		defer stop1()
		next2, stop2 := iter.Pull(s2)
		defer stop2()

		// Get initial values
		v1, ok1 := next1()
		v2, ok2 := next2()

		// Main merge loop
		for ok1 || ok2 {
			if !ok2 || (ok1 && ((v1 <= v2) != desc)) {
				if !yield(v1) {
					return
				}
				v1, ok1 = next1()
			} else {
				if !yield(v2) {
					return
				}
				v2, ok2 = next2()
			}
		}
	}
}

// MergeOrdered2 returns an iterator that yields key-value pairs from both input sequences
// in ascending or descending order based on desc flag and the value component.
func MergeOrdered2[K any, V cmp.Ordered](s1, s2 iter.Seq2[K, V], desc bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next1, stop1 := iter.Pull2(s1)
		defer stop1()
		next2, stop2 := iter.Pull2(s2)
		defer stop2()

		// Get initial values
		k1, v1, ok1 := next1()
		k2, v2, ok2 := next2()

		// Main merge loop
		for ok1 || ok2 {
			if !ok2 || (ok1 && ((v1 <= v2) != desc)) {
				if !yield(k1, v1) {
					return
				}
				k1, v1, ok1 = next1()
			} else {
				if !yield(k2, v2) {
					return
				}
				k2, v2, ok2 = next2()
			}
		}
	}
}
