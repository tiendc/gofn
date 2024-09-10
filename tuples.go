package gofn

type Tuple2[T1, T2 any] struct {
	Elem1 T1
	Elem2 T2
}

// Unpack unpacks elements of tuple, panic on nil pointer
func (tuple2 *Tuple2[T1, T2]) Unpack() (T1, T2) {
	return tuple2.Elem1, tuple2.Elem2
}

type Tuple3[T1, T2, T3 any] struct {
	Elem1 T1
	Elem2 T2
	Elem3 T3
}

// Unpack unpacks elements of tuple, panic on nil pointer
func (tuple3 *Tuple3[T1, T2, T3]) Unpack() (T1, T2, T3) {
	return tuple3.Elem1, tuple3.Elem2, tuple3.Elem3
}

type Tuple4[T1, T2, T3, T4 any] struct {
	Elem1 T1
	Elem2 T2
	Elem3 T3
	Elem4 T4
}

// Unpack unpacks elements of tuple, panic on nil pointer
func (tuple4 *Tuple4[T1, T2, T3, T4]) Unpack() (T1, T2, T3, T4) {
	return tuple4.Elem1, tuple4.Elem2, tuple4.Elem3, tuple4.Elem4
}

type Tuple5[T1, T2, T3, T4, T5 any] struct {
	Elem1 T1
	Elem2 T2
	Elem3 T3
	Elem4 T4
	Elem5 T5
}

// Unpack unpacks elements of tuple, panic on nil pointer
func (tuple5 *Tuple5[T1, T2, T3, T4, T5]) Unpack() (T1, T2, T3, T4, T5) {
	return tuple5.Elem1, tuple5.Elem2, tuple5.Elem3, tuple5.Elem4, tuple5.Elem5
}
