package returns

//go:nosplit
//go:noinline
func ExplicitReturns(a int, b int) (c int, d int) {
	c = a / b
	d = a % b
	return c, d
}

//go:nosplit
//go:noinline
func ImplicitReturns(a int, b int) (c int, d int) {
	c = a / b
	d = a % b
	return
}
