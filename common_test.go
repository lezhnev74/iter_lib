package iter_lib_test

import "iter"

type kvGroup struct {
	h  string
	kv []kv
}

type kv struct {
	k string
	v int
}

func seq2ToKvs(s iter.Seq2[string, int]) []kv {
	var r []kv
	for k, v := range s {
		r = append(r, kv{k, v})
	}
	return r
}
func kvsToSeq(s []kv) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for _, kv := range s {
			if !yield(kv.k, kv.v) {
				return
			}
		}
	}
}
