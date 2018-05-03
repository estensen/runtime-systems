package main

import (
	"github.com/estensen/runtime-systems/benchmarks/hello"

	"github.com/estensen/runtime-systems/benchmarks/profiler"
)

func main() {
	profiler.Profiler()
	hello.Hello()
}
