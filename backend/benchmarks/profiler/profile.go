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

func Profiler(packageName string, profilingDone chan bool) {
	runPackage(packageName, profilingDone)
}

func runPackage(packageName string, profilingDone chan bool) {
	defer profilingIsDone(profilingDone)
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	switch packageName {
	case "sort":
		sort.Sort()
	case "fiboncci":
		fibonacci.Fibonacci(100000000)
	case "wordSearch":
		//change WordSearchWithList to WordSearchWithMap to see difference in CPU time
		wordSearch.WordSearchWithList("brick")
	default:
		fmt.Println("Package not found")
	}
}

func profilingIsDone(profilingDone chan bool) {
	profilingDone <- true
}

func CPUPercent() []string {
	c, _ := cpu.Percent(45*time.Millisecond, false)
	t := time.Now()
	now := t.Format("15:04:05")
	line := []string{now, strconv.FormatFloat(c[0], 'f', 2, 64)} // Convert c from float64 to string
	return line
}
