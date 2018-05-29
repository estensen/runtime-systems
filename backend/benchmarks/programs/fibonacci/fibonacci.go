package fibonacci

import (
	"fmt"
)

func Fibonacci(n int) {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	fmt.Println(b)
}
