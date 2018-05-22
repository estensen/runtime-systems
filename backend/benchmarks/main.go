package main

import (
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		print("Please run go with a given program. Example: 'go run main.go fibonacci')")
	} else {
		//packageName := os.Args[1]
		//profiler.Profiler(packageName)
		//cpuTextfileToJSON(packageName)
	}
}

func cpuTextfileToJSON(packageName string) {
	filename := packageName + ".txt"

	//create textfile to save terminal output in. File is created en reports directory
	file, err := os.Create("../reports/" + filename)
	if err != nil {
		panic("Could not create " + filename)
	}
	defer file.Close()

	//run command to create text from pprof
	pproftext := exec.Command("go", "tool", "pprof", "-text", "cpu.pprof")
	textOut, err := pproftext.Output()
	if err != nil {
		panic(err)
	}

	//save command output in textfile
	file.Write(textOut)
}
