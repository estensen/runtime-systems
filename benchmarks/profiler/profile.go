package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/pkg/profile"
)

func fibonacci(n int) {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, a+b

		fmt.Println(b)
	}
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	fibonacci(1000000)
}
