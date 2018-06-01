package profiler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/estensen/runtime-systems/backend/benchmarks/programs/fibonacci"
	"github.com/estensen/runtime-systems/backend/benchmarks/programs/sort"
	"github.com/estensen/runtime-systems/backend/benchmarks/programs/wordSearch"
	"github.com/pkg/profile"
	"github.com/shirou/gopsutil/cpu"
)

func Profiler(packageName string, profilingRunning chan bool) {
	runPackage(packageName, profilingRunning)
}

func runPackage(packageName string, profilingRunning chan bool) {
	defer profilingIsRunning(profilingRunning)
	defer profile.Start(profile.CPUProfile, profile.ProfilePath("./benchmarks/programs/"+packageName)).Stop()
	switch packageName {
	case "sort":
		sort.Sort()
	case "fibonacci":
		fibonacci.Fibonacci(100000000)
	case "wordSearch":
		//change WordSearchWithList to WordSearchWithMap to see difference in CPU time
		wordSearch.WordSearchWithList("brick")
	default:
		fmt.Println("Package not found")
	}
}

func profilingIsRunning(profilingRunning chan bool) {
	profilingRunning <- false
}

func CPUPercent() []string {
	c, _ := cpu.Percent(45*time.Millisecond, false)
	t := time.Now()
	now := t.Format("15:04:05")
	line := []string{now, strconv.FormatFloat(c[0], 'f', 2, 64)} // Convert c from float64 to string
	return line
}
