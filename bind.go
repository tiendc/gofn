package gofn

func Bind1Arg0Ret[A1 any](fn func(A1), a1 A1) func() {
	return func() {
		fn(a1)
	}
}

func Bind1Arg1Ret[A1 any, R1 any](fn func(A1) R1, a1 A1) func() R1 {
	return func() R1 {
		return fn(a1)
	}
}

func Bind1Arg2Ret[A1 any, R1 any, R2 any](fn func(A1) (R1, R2), a1 A1) func() (R1, R2) {
	return func() (R1, R2) {
		return fn(a1)
	}
}

func Bind1Arg3Ret[A1 any, R1 any, R2 any, R3 any](fn func(A1) (R1, R2, R3), a1 A1) func() (R1, R2, R3) {
	return func() (R1, R2, R3) {
		return fn(a1)
	}
}

func Bind2Arg0Ret[A1 any, A2 any](fn func(A1, A2), a1 A1, a2 A2) func() {
	return func() {
		fn(a1, a2)
	}
}

func Bind2Arg1Ret[A1 any, A2 any, R1 any](fn func(A1, A2) R1, a1 A1, a2 A2) func() R1 {
	return func() R1 {
		return fn(a1, a2)
	}
}

func Bind2Arg2Ret[A1 any, A2 any, R1 any, R2 any](fn func(A1, A2) (R1, R2), a1 A1, a2 A2) func() (R1, R2) {
	return func() (R1, R2) {
		return fn(a1, a2)
	}
}

// nolint: lll
func Bind2Arg3Ret[A1 any, A2 any, R1 any, R2 any, R3 any](fn func(A1, A2) (R1, R2, R3), a1 A1, a2 A2) func() (R1, R2, R3) {
	return func() (R1, R2, R3) {
		return fn(a1, a2)
	}
}

func Bind3Arg0Ret[A1 any, A2 any, A3 any](fn func(A1, A2, A3), a1 A1, a2 A2, a3 A3) func() {
	return func() {
		fn(a1, a2, a3)
	}
}

func Bind3Arg1Ret[A1 any, A2 any, A3 any, R1 any](fn func(A1, A2, A3) R1, a1 A1, a2 A2, a3 A3) func() R1 {
	return func() R1 {
		return fn(a1, a2, a3)
	}
}

// nolint: lll
func Bind3Arg2Ret[A1 any, A2 any, A3 any, R1 any, R2 any](fn func(A1, A2, A3) (R1, R2), a1 A1, a2 A2, a3 A3) func() (R1, R2) {
	return func() (R1, R2) {
		return fn(a1, a2, a3)
	}
}

// nolint: lll
func Bind3Arg3Ret[A1 any, A2 any, A3 any, R1 any, R2 any, R3 any](fn func(A1, A2, A3) (R1, R2, R3), a1 A1, a2 A2, a3 A3) func() (R1, R2, R3) {
	return func() (R1, R2, R3) {
		return fn(a1, a2, a3)
	}
}

func Bind4Arg0Ret[A1 any, A2 any, A3 any, A4 any](fn func(A1, A2, A3, A4), a1 A1, a2 A2, a3 A3, a4 A4) func() {
	return func() {
		fn(a1, a2, a3, a4)
	}
}

// nolint: lll
func Bind4Arg1Ret[A1 any, A2 any, A3 any, A4 any, R1 any](fn func(A1, A2, A3, A4) R1, a1 A1, a2 A2, a3 A3, a4 A4) func() R1 {
	return func() R1 {
		return fn(a1, a2, a3, a4)
	}
}

// nolint: lll
func Bind4Arg2Ret[A1 any, A2 any, A3 any, A4 any, R1 any, R2 any](fn func(A1, A2, A3, A4) (R1, R2), a1 A1, a2 A2, a3 A3, a4 A4) func() (R1, R2) {
	return func() (R1, R2) {
		return fn(a1, a2, a3, a4)
	}
}

// nolint: lll
func Bind4Arg3Ret[A1 any, A2 any, A3 any, A4 any, R1 any, R2 any, R3 any](fn func(A1, A2, A3, A4) (R1, R2, R3), a1 A1, a2 A2, a3 A3, a4 A4) func() (R1, R2, R3) {
	return func() (R1, R2, R3) {
		return fn(a1, a2, a3, a4)
	}
}

// nolint: lll
func Bind5Arg0Ret[A1 any, A2 any, A3 any, A4 any, A5 any](fn func(A1, A2, A3, A4, A5), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) func() {
	return func() {
		fn(a1, a2, a3, a4, a5)
	}
}

// nolint: lll
func Bind5Arg1Ret[A1 any, A2 any, A3 any, A4 any, A5 any, R1 any](fn func(A1, A2, A3, A4, A5) R1, a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) func() R1 {
	return func() R1 {
		return fn(a1, a2, a3, a4, a5)
	}
}

// nolint: lll
func Bind5Arg2Ret[A1 any, A2 any, A3 any, A4 any, A5 any, R1 any, R2 any](fn func(A1, A2, A3, A4, A5) (R1, R2), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) func() (R1, R2) {
	return func() (R1, R2) {
		return fn(a1, a2, a3, a4, a5)
	}
}

// nolint: lll
func Bind5Arg3Ret[A1 any, A2 any, A3 any, A4 any, A5 any, R1 any, R2 any, R3 any](fn func(A1, A2, A3, A4, A5) (R1, R2, R3), a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) func() (R1, R2, R3) {
	return func() (R1, R2, R3) {
		return fn(a1, a2, a3, a4, a5)
	}
}

func Partial2Arg0Ret[A1 any, A2 any](fn func(A1, A2), a1 A1) func(A2) {
	return func(a2 A2) {
		fn(a1, a2)
	}
}

func Partial2Arg1Ret[A1 any, A2 any, R1 any](fn func(A1, A2) R1, a1 A1) func(A2) R1 {
	return func(a2 A2) R1 {
		return fn(a1, a2)
	}
}

func Partial2Arg2Ret[A1 any, A2 any, R1 any, R2 any](fn func(A1, A2) (R1, R2), a1 A1) func(A2) (R1, R2) {
	return func(a2 A2) (R1, R2) {
		return fn(a1, a2)
	}
}

// nolint: lll
func Partial2Arg3Ret[A1 any, A2 any, R1 any, R2 any, R3 any](fn func(A1, A2) (R1, R2, R3), a1 A1) func(A2) (R1, R2, R3) {
	return func(a2 A2) (R1, R2, R3) {
		return fn(a1, a2)
	}
}

func Partial3Arg0Ret[A1 any, A2 any, A3 any](fn func(A1, A2, A3), a1 A1) func(A2, A3) {
	return func(a2 A2, a3 A3) {
		fn(a1, a2, a3)
	}
}

func Partial3Arg1Ret[A1 any, A2 any, A3 any, R1 any](fn func(A1, A2, A3) R1, a1 A1) func(A2, A3) R1 {
	return func(a2 A2, a3 A3) R1 {
		return fn(a1, a2, a3)
	}
}

// nolint: lll
func Partial3Arg2Ret[A1 any, A2 any, A3 any, R1 any, R2 any](fn func(A1, A2, A3) (R1, R2), a1 A1) func(A2, A3) (R1, R2) {
	return func(a2 A2, a3 A3) (R1, R2) {
		return fn(a1, a2, a3)
	}
}

// nolint: lll
func Partial3Arg3Ret[A1 any, A2 any, A3 any, R1 any, R2 any, R3 any](fn func(A1, A2, A3) (R1, R2, R3), a1 A1) func(A2, A3) (R1, R2, R3) {
	return func(a2 A2, a3 A3) (R1, R2, R3) {
		return fn(a1, a2, a3)
	}
}

func Partial4Arg0Ret[A1 any, A2 any, A3 any, A4 any](fn func(A1, A2, A3, A4), a1 A1) func(A2, A3, A4) {
	return func(a2 A2, a3 A3, a4 A4) {
		fn(a1, a2, a3, a4)
	}
}

// nolint: lll
func Partial4Arg1Ret[A1 any, A2 any, A3 any, A4 any, R1 any](fn func(A1, A2, A3, A4) R1, a1 A1) func(A2, A3, A4) R1 {
	return func(a2 A2, a3 A3, a4 A4) R1 {
		return fn(a1, a2, a3, a4)
	}
}

// nolint: lll
func Partial4Arg2Ret[A1 any, A2 any, A3 any, A4 any, R1 any, R2 any](fn func(A1, A2, A3, A4) (R1, R2), a1 A1) func(A2, A3, A4) (R1, R2) {
	return func(a2 A2, a3 A3, a4 A4) (R1, R2) {
		return fn(a1, a2, a3, a4)
	}
}

// nolint: lll
func Partial4Arg3Ret[A1 any, A2 any, A3 any, A4 any, R1 any, R2 any, R3 any](fn func(A1, A2, A3, A4) (R1, R2, R3), a1 A1) func(A2, A3, A4) (R1, R2, R3) {
	return func(a2 A2, a3 A3, a4 A4) (R1, R2, R3) {
		return fn(a1, a2, a3, a4)
	}
}

// nolint: lll
func Partial5Arg0Ret[A1 any, A2 any, A3 any, A4 any, A5 any](fn func(A1, A2, A3, A4, A5), a1 A1) func(A2, A3, A4, A5) {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) {
		fn(a1, a2, a3, a4, a5)
	}
}

// nolint: lll
func Partial5Arg1Ret[A1 any, A2 any, A3 any, A4 any, A5 any, R1 any](fn func(A1, A2, A3, A4, A5) R1, a1 A1) func(A2, A3, A4, A5) R1 {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) R1 {
		return fn(a1, a2, a3, a4, a5)
	}
}

// nolint: lll
func Partial5Arg2Ret[A1 any, A2 any, A3 any, A4 any, A5 any, R1 any, R2 any](fn func(A1, A2, A3, A4, A5) (R1, R2), a1 A1) func(A2, A3, A4, A5) (R1, R2) {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2) {
		return fn(a1, a2, a3, a4, a5)
	}
}

// nolint: lll
func Partial5Arg3Ret[A1 any, A2 any, A3 any, A4 any, A5 any, R1 any, R2 any, R3 any](fn func(A1, A2, A3, A4, A5) (R1, R2, R3), a1 A1) func(A2, A3, A4, A5) (R1, R2, R3) {
	return func(a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2, R3) {
		return fn(a1, a2, a3, a4, a5)
	}
}
