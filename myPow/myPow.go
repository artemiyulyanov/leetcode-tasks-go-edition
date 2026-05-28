package myPow

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return MyPow(1/x, -n)
	}

	if n%2 == 0 {
		return MyPow(x*x, n/2)
	}

	return x * MyPow(x, n-1)
}
