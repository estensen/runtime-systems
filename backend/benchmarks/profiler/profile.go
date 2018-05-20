package profiler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/estensen/runtime-systems/backend/benchmarks/programs/sort"
	"github.com/pkg/profile"
	"github.com/shirou/gopsutil/cpu"
)

//Profiler running from main
func Profiler(packageName string, profilingDone chan bool) {
	runPackage(packageName, profilingDone)
}

func runPackage(packageName string, profilingDone chan bool) {
	//checkPackage
	//checkPackageTest
	//go CPUPercent()
	defer profilingIsDone(profilingDone)
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	switch packageName {
	case "sort":
		sort.Sort()
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
	now := t.Format("2006-01-02 15:04:05")
	line := []string{now, strconv.FormatFloat(c[0], 'f', 2, 64)} // Convert c from float64 to string
	return line

	/*
		filename := "cpupercent.csv"

		c, _ := cpu.Percent(10*time.Millisecond, false)
		t := time.Now()
		now := t.Format("2006-01-02 15:04:05")

		// Write file as <timestamp>,<size>,<board>
		line := []string{now, strconv.FormatFloat(c[0], 'f', 2, 64)} // Convert c from float64 to string

		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
		if err != nil {
			panic("could not open csv file")
		}

		defer file.Close()

		json, err = json.Marshal(line)
		if err != nil {
			panic(err)
		}

		file.Write(json)*/
}
