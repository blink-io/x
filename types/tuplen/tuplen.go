package tuplen

type Tuple2[V1, V2 any] struct {
	V1 V1
	V2 V2
}

type Tuple3[V1, V2, V3 any] struct {
	V1 V1
	V2 V2
	V3 V3
}

type Tuple4[V1, V2, V3, V4 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
}

type Tuple5[V1, V2, V3, V4, V5 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
}

type Tuple6[V1, V2, V3, V4, V5, V6 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
}

type Tuple7[V1, V2, V3, V4, V5, V6, V7 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
	V7 V7
}

type Tuple8[V1, V2, V3, V4, V5, V6, V7, V8 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
	V7 V7
	V8 V8
}

type Tuple9[V1, V2, V3, V4, V5, V6, V7, V8, V9 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
	V5 V5
	V6 V6
	V7 V7
	V8 V8
	V9 V9
}

func Of2[V1, V2 any](v1 V1, v2 V2) Tuple2[V1, V2] {
	return Tuple2[V1, V2]{
		V1: v1,
		V2: v2,
	}
}

func Of3[V1, V2, V3 any](v1 V1, v2 V2, v3 V3) Tuple3[V1, V2, V3] {
	return Tuple3[V1, V2, V3]{
		V1: v1,
		V2: v2,
		V3: v3,
	}
}

func Of4[V1, V2, V3, V4 any](v1 V1, v2 V2, v3 V3, v4 V4) Tuple4[V1, V2, V3, V4] {
	return Tuple4[V1, V2, V3, V4]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
	}
}
