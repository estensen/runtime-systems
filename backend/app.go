package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/estensen/runtime-systems/backend/benchmarks/profiler"
	"github.com/julienschmidt/httprouter"
)

var graphPoints = make(map[string][]string)
var profilingStarted = false
var profilingDoneChannel = make(chan bool, 1)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	payload := map[string]string{"apiType": "This is a RESTful API"}
	respondWithJSON(w, http.StatusOK, payload)
}

func getPrograms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dirname := "benchmarks/programs"

	f, err := os.Open(dirname)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Could not open program files")
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Could not read program files")
	}

	var programs []string

	for _, program := range files {
		programs = append(programs, program.Name())
	}

	payload := map[string][]string{"programs": programs}

	respondWithJSON(w, http.StatusOK, payload)
}

func runProfiling(packageName string) {
	profilingisDone := false

	if !profilingStarted {
		profilingStarted = true
		go func() {
			profiler.Profiler(packageName, profilingDoneChannel)
		}()

		for !profilingisDone {
			cpuStats := profiler.CPUPercent()
			graphPoints["Time"] = append(graphPoints["Time"], cpuStats[0])
			graphPoints["Percent"] = append(graphPoints["Percent"], cpuStats[1])

			timer := time.NewTimer(50 * time.Millisecond)
			<-timer.C
			profilingisDone = checkIfProfilingisDone()
		}
	}
}

func checkIfPprofFileExists() bool {
	if _, err := os.Stat("cpu.pprof"); os.IsNotExist(err) {
		fmt.Println("pprof does not exists")
		return false
	}
	fmt.Println("pprof exists")
	return true
}

func checkIfDiagramExists(packageName string) bool {
	if _, err := os.Stat("./diagrams/" + packageName); os.IsNotExist(err) {
		fmt.Println("diagram does not exists")
		return false
	}
	fmt.Println("diagram exists")
	return true
}

func checkIfReportExists(packageName string) bool {
	if _, err := os.Stat("./reports/" + packageName); os.IsNotExist(err) {
		fmt.Println("report does not exists")
		return false
	}
	fmt.Println("report exists")
	return true
}

// Read file and return file length
func getCPUreport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("program")
	filename := packageName + ".txt"

	if !checkIfPprofFileExists() {
		runProfiling(packageName)
	}
	if checkIfReportExists(filename) {
		report, err := os.Open("reports/" + filename)
		if err != nil {
			panic("Could not open " + filename)
		}
		respondWithJSON(w, http.StatusOK, report)
	} else {
		//create textfile to save terminal output in. File is created en reports directory
		report, err := os.Create("reports/" + filename)
		if err != nil {
			panic("Could not create " + filename)
		}
		defer report.Close()

		//run command to create text from pprof
		pproftext := exec.Command("go", "tool", "pprof", "-text", "cpu.pprof")
		reportText, err := pproftext.Output()
		if err != nil {
			panic(err)
		}

		//save command output in textfile
		report.Write(reportText)
		
		respondWithJSON(w, http.StatusOK, reportText)
	}
}

// This method will run our profiler with the given package name
// then create a PDF diagram with the given cpu.pprof file and show this on the webpage.
func getCPUdiagram(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("program")
	filename := packageName + ".png"

	if !checkIfPprofFileExists() {
		runProfiling(packageName)
	}

	if checkIfDiagramExists(filename) {
		diagram, err := ioutil.ReadFile("diagrams/" + filename)
		if err != nil {
			panic("Could not open " + filename)
		}
		respondWithPNG(w, http.StatusOK, diagram)
	} else {
		//create textfile to save terminal output in. File is created in reports directory
		diagram, err := os.Create("diagrams/" + filename)
		if err != nil {
			panic("Could not create " + filename)
		}
		defer diagram.Close()

		//run command to create text from pprof
		pprofPNG := exec.Command("go", "tool", "pprof", "-png", "cpu.pprof")
		png, err := pprofPNG.Output()
		if err != nil {
			panic(err)
		}

		diagram.Write(png)
		respondWithPNG(w, http.StatusOK, png)
	}
}

func getGraphData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("program")
	if !checkIfPprofFileExists() {
		runProfiling(packageName)
	}
	respondWithJSON(w, http.StatusOK, graphPoints)
}

func checkIfProfilingisDone() bool {
	select {
	case done := <-profilingDoneChannel:
		fmt.Println("profiling is Done")
		return done
	default:
		return false
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithPNG(w http.ResponseWriter, code int, payload []byte) {
	enableCors(&w)

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(code)
	w.Write(payload)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/cpu", getPrograms)
	router.GET("/cpu/report/:program", getCPUreport)
	router.GET("/cpu/diagram/:program", getCPUdiagram)
	router.GET("/cpu/graph/:program", getGraphData)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running API server in prod mode")
	} else {
		log.Println("Running API server in dev mode")
	}

	http.ListenAndServe("localhost:8080", router)
}
