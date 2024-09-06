package pair

type Pair[L, R any] struct {
	L L
	R R
}

func Of[L, R any](l L, r R) Pair[L, R] {
	return Pair[L, R]{L: l, R: r}
}
