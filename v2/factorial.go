package arithmetic

func Factorial(n int) int {
	var f int = 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

// new function in v2

func IsEven(n int) bool {
	return n%2 == 0
}
