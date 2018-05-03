package main

import (
	_ "net/http/pprof"

	"github.com/estensen/runtime-systems/benchmarks/fibonacci"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	fibonacci.Fibonacci(1000000)
}
