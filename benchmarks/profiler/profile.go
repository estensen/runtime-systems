package profiler

import (
	"fmt"

	"github.com/estensen/runtime-systems/benchmarks/fibonacci"
	"github.com/pkg/profile"
)

//Profiler running from main
func Profiler() {
	fmt.Println("Running Fibonacci Profiler")
	runFibonacciProfiler()
}

func runFibonacciProfiler() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	fibonacci.Fibonacci(1000000)
}
