package algorithm

func djb2(s string) uint64 {
	var hash uint64 = 5381
	for c := range s {
		hash = hash * 33 + uint64(c)
	}
	return hash
}
