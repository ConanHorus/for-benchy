package fib

var fibMap = make(map[int]int)

func Fibonacci(n int) int {
	if n <= 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciWithCache(n int) int {
	if n <= 1 {
		return 1
	}

	if value, ok := fibMap[n]; ok {
		return value
	}

	value := FibonacciWithCache(n-1) + FibonacciWithCache(n-2)
	fibMap[n] = value
	return value
}
