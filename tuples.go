package gofn

type Tuple2[T any, U any] struct {
	Elem1 T
	Elem2 U
}

type Tuple3[T any, U any, V any] struct {
	Elem1 T
	Elem2 U
	Elem3 V
}
