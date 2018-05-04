package main

import (
	"os"

	"github.com/estensen/runtime-systems/benchmarks/profiler"
)

func main() {
	if len(os.Args) != 2 {
		print("Please run go with a given program. Example: 'go run main.go fibonacci')")
	} else {
		packageName := os.Args[1]
		//fmt.Println(packageName)
		profiler.Profiler(packageName)
	}
}
