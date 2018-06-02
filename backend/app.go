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
var profilingDoneChannel = make(chan bool, 1)
var profilingIsRunning = false

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

func deleteOldProfile(program string, profileType string) {
	programPath := "benchmarks/programs/" + program + "/"
	pprofPath := programPath + profileType + ".pprof"

	//delete pprof
	_, err := os.Open(pprofPath)
	if err == nil {
		os.Remove(pprofPath)
	}

	//delete diagrams
	diagramPath := "diagrams"
	emptyFolder(diagramPath, program, profileType, ".png")

	//delete reports
	reportPath := "reports"
	emptyFolder(reportPath, program, profileType, ".txt")

}

func emptyFolder(dir string, program string, profileType string, format string) {
	directory, err := os.Open(dir)
	if err != nil {
		panic("unable to Open " + dir)
	}

	dirFiles, err := directory.Readdir(0)
	if err != nil {
		panic("unable to read dirfiles ")
	}

	for index := range dirFiles {
		file := dirFiles[index]
		name := file.Name()

		programName := program + "_" + profileType + format

		if name == programName {
			fullpath := dir + "/" + name

			os.Remove(fullpath)
		}
	}
}

func runProfiling(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("program")
	profileType := ps.ByName("profileType")

	if !profilingIsRunning {
		profilingIsRunning = true
		deleteOldProfile(packageName, profileType)
		go func() {
			profiler.Profiler(packageName, profileType, profilingDoneChannel)
		}()

		for profilingIsRunning {
			cpuStats := profiler.CPUPercent()
			graphPoints["Time"] = append(graphPoints["Time"], cpuStats[0])
			graphPoints["Percent"] = append(graphPoints["Percent"], cpuStats[1])

			timer := time.NewTimer(50 * time.Millisecond)
			<-timer.C
			profilingIsRunning = checkIfProfilingisDone()
		}
	}
	payload := map[string]bool{"isProfiled": true} // Hardcoded
	respondWithJSON(w, http.StatusOK, payload)
}

func checkIfPprofFileExists(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("program")
	profileType := ps.ByName("profileType")
	if profileType == "memory" {
		profileType = "mem"
	}

	payload := map[string]bool{"profileExists": true}
	pprofPath := "benchmarks/programs/" + packageName + "/" + profileType + ".pprof"
	fmt.Println(pprofPath)
	if _, err := os.Stat(pprofPath); os.IsNotExist(err) {
		payload["profileExists"] = false
	}
	respondWithJSON(w, http.StatusOK, payload)
}

func checkIfReportExists(packageName string) bool {
	if _, err := os.Stat("./reports/" + packageName); os.IsNotExist(err) {
		return false
	}
	return true
}

func createReport(packageName string, profileType string) {
	file, err := os.Create("reports/" + packageName + "_" + profileType + ".txt")
	if err != nil {
		panic("Could not create textfile" + packageName)
	}
	defer file.Close()

	if profileType == "memory" {
		profileType = "mem"
	}

	pprofPath := "benchmarks/programs/" + packageName + "/" + profileType + ".pprof"

	pproftext := exec.Command("go", "tool", "pprof", "-text", pprofPath)
	reportText, err := pproftext.Output()
	if err != nil {
		panic(err)
	}

	file.Write(reportText)
}

func readReport(filename string) string {
	report, err := ioutil.ReadFile("reports/" + filename)
	if err != nil {
		panic("Could not open " + filename)
	}

	reportStr := string(report)

	return reportStr
}

// Read file and return file length
func getReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("program")
	profileType := ps.ByName("profileType")
	filename := packageName + "_" + profileType + ".txt"

	if !checkIfReportExists(filename) {
		createReport(packageName, profileType)
	}

	reportStr := readReport(filename)

	respondWithJSON(w, http.StatusOK, reportStr)
}

// This method will run our profiler with the given package name
// then create a PDF diagram with the given cpu.pprof file and show this on the webpage.
func getDiagram(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("program")
	profileType := ps.ByName("profileType")
	filename := packageName + "_" + profileType + ".png"

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

		if profileType == "memory" {
			profileType = "mem"
		}

		//run command to create text from pprof
		pprofPath := "benchmarks/programs/" + packageName + "/" + profileType + ".pprof"
		pprofPNG := exec.Command("go", "tool", "pprof", "-png", pprofPath)
		png, err := pprofPNG.Output()
		if err != nil {
			panic(err)
		}

		diagram.Write(png)
		respondWithPNG(w, http.StatusOK, png)
	}
}

func checkIfDiagramExists(packageName string) bool {
	if _, err := os.Stat("./diagrams/" + packageName); os.IsNotExist(err) {
		return false
	}
	return true
}

func getGraphData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	respondWithJSON(w, http.StatusOK, graphPoints)
}

func checkIfProfilingisDone() bool {
	select {
	case done := <-profilingDoneChannel:
		fmt.Println("profiling is Done")
		return done
	default:
		return true
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
	router.GET("/programs/:profileType", getPrograms)
	router.GET("/checkProfiling/:profileType/:program", checkIfPprofFileExists)
	router.GET("/report/:profileType/:program", getReport)
	router.GET("/diagram/:profileType/:program", getDiagram)
	router.GET("/graph/:profileType/:program", getGraphData)
	router.GET("/runprofiling/:profileType/:program", runProfiling)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running API server in prod mode")
	} else {
		log.Println("Running API server in dev mode")
	}

	http.ListenAndServe("localhost:8080", router)
}
