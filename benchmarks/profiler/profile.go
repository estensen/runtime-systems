package profiler

import (
	"fmt"

	"github.com/estensen/runtime-systems/benchmarks/sort"

	"github.com/pkg/profile"
)

//Profiler running from main
func Profiler(packageName string) {
	runPackage(packageName)
}

func runPackage(packageName string) {
	//checkPackage
	//checkPackageTest
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	switch packageName {
	case "sort":
		sort.Sort()
	default:
		fmt.Println("Package not found")
	}
}
