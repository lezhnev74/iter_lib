package iter_lib

import "iter"

// Map returns an iterator that yields the values from s transformed by f.
func Map[V any](seq iter.Seq[V], f func(V) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Map2 returns an iterator that yields the values from s transformed by f.
// The values are paired with their corresponding keys.
func Map2[K, V, R any](s iter.Seq2[K, V], f func(V) R) iter.Seq2[K, R] {
	return func(yield func(K, R) bool) {
		for k, v := range s {
			if !yield(k, f(v)) {
				return
			}
		}
	}
}
