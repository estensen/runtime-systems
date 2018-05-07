package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/estensen/runtime-systems/backend/benchmarks/profiler"
	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	payload := map[string]string{"apiType": "This is a RESTful API"}
	respondWithJSON(w, http.StatusOK, payload)
}

func getReportsCPU(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dirname := "reports"

	f, err := os.Open(dirname)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Could not open report files")
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Could not read report files")
	}

	var filenames []string

	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	payload := map[string][]string{"filenames": filenames}

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

func getCPUdiagram(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("package")
	fmt.Println(packageName)

	profiler.Profiler(packageName)
	/*
		run terminal-commando go tool pprof to save in textfile in diagrams directory
		if err := exec.Command("go", "tool", "pprof", "-pdf", "cpu.pprof").Run(); err != nil {
			respondWithJSON(w, http.StatusNotFound, "Unable to run command line")
		}
			if err := exec.Command("exit").Run(); err != nil {
				respondWithJSON(w, http.StatusNotFound, "Unable to close command pprof program")
			}
	*/

	dirname := "diagrams/"
	filename := dirname + packageName + "Example.png"

	pdf, err := ioutil.ReadFile(filename)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Could not read file - " + http.StatusText(404)))
	} else {
		w.Write(pdf)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/cpu", getReportsCPU)
	router.GET("/cpu/reports/:filename", getReportCPU)
	router.GET("/cpu/diagram/:package", getCPUdiagram)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running API server in prod mode")
	} else {
		log.Println("Running API server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}
