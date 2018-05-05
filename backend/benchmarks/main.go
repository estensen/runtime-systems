package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/estensen/runtime-systems/benchmarks/profiler"
)

func main() {
	if len(os.Args) != 2 {
		print("Please run go with a given program. Example: 'go run main.go fibonacci')")
	} else {
		packageName := os.Args[1]
		profiler.Profiler(packageName)
		cpuTextfileToJSON(packageName)
	}
}

func cpuTextfileToJSON(name string) {
	filename := name + ".txt"

	file, err := os.Open(filename)
	if err != nil {
		panic("Could not open " + filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
