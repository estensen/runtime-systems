package profiler

import (
	"fmt"
	"os"
	"time"

	"github.com/estensen/runtime-systems/backend/benchmarks/programs/sort"
	"github.com/pkg/profile"
	"github.com/shirou/gopsutil/cpu"
)

//Profiler running from main
func Profiler(packageName string) {
	runPackage(packageName)
}

func runPackage(packageName string) {
	//checkPackage
	//checkPackageTest
	go CPUPercent()
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	switch packageName {
	case "sort":
		sort.Sort()
	default:
		fmt.Println("Package not found")
	}
}

func CPUPercent() {
	filename := "cpupercent.csv"

	c, _ := cpu.Percent(500*time.Millisecond, false)

	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")

	// Write file as <timestamp>,<size>,<board>
	line := []string{now, c} // Convert c from float64 to string

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		panic("could not open csv file")
	}

	defer file.Close()

}
