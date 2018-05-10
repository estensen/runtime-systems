package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

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

// When going to localhost:8080/cpu/diagram/:package, this method will run our profiler with the given package name
// then create a PDF diagram with the given cpu.pprof file and show this on the webpage.
func getCPUdiagram(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	packageName := ps.ByName("package")
	profiler.Profiler(packageName)

	filename := packageName + ".png"

	//create textfile to save terminal output in. File is created en reports directory
	file, err := os.Create("diagrams/" + filename)
	if err != nil {
		panic("Could not create " + filename)
	}
	defer file.Close()

	//run command to create text from pprof
	pprofPDF := exec.Command("go", "tool", "pprof", "-png", "cpu.pprof")
	pdf, err := pprofPDF.Output()
	if err != nil {
		panic(err)
	}

	//save command output in textfile
	file.Write(pdf)

	//Write PDF to site
	w.Write(pdf)

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
