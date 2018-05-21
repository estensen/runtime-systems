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

/*
type point struct {
	Time    string
	Percent string
}*/

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

// Read file and return file length
func getReportCPU(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dirname := "reports/"
	filename := dirname + ps.ByName("filename")

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		respondWithJSON(w, http.StatusNotFound, "File not found")
	} else {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			respondWithJSON(w, http.StatusBadRequest, "Could not read report file")
		}

		reportLength := len(file)

		payload := map[string]int{"reportLength": reportLength}
		respondWithJSON(w, http.StatusOK, payload)
	}
}

// This method will run our profiler with the given package name
// then create a PDF diagram with the given cpu.pprof file and show this on the webpage.
func getCPUdiagram(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("package")
	profiler.Profiler(packageName, profilingDoneChannel)

	filename := packageName + ".png"

	//create textfile to save terminal output in. File is created en reports directory
	file, err := os.Create("diagrams/" + filename)
	if err != nil {
		panic("Could not create " + filename)
	}
	defer file.Close()

	//run command to create text from pprof
	pprofPNG := exec.Command("go", "tool", "pprof", "-png", "cpu.pprof")
	png, err := pprofPNG.Output()
	if err != nil {
		panic(err)
	}

	file.Write(png)
	respondWithPNG(w, http.StatusOK, png)
}

func getLiveData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("package")
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
	router.GET("/cpu/reports/:filename", getReportCPU)
	router.GET("/cpu/diagram/:package", getCPUdiagram)
	router.GET("/cpu/live/:package", getLiveData)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running API server in prod mode")
	} else {
		log.Println("Running API server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}
