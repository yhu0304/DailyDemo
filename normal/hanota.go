package normal

func hanota(A []int, B []int, C []int) []int {
	if len(A) != 0 {
		A, B, C = move(len(A), A, B, C)
	}
	return C
}

func move(n int, A, B, C []int) (resA, resB, resC []int) {
	if n > 1 {
		A, C, B = move(n-1, A, C, B)
	}
	C = append(C, A[len(A)-1])
	A = A[:len(A)-1]
	if n > 1 {
		B, A, C = move(n-1, B, A, C)
	}
	return A, B, C
}
