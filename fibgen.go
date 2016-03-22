package go2test

// NewFibGen returns a Fibonacci generator
func NewFibGen() func() int {
	a, b := 0, 1
	return func() int {
		a, b = a+b, a
		return b
	}
}
