package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	filename := "reports/" + ps.ByName("filename")
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Coul not read report file")
	}

	fmt.Printf("\nFilename: %s", filename)
	fmt.Printf("\nFile: %s", file)
	fmt.Printf("\nLength: %d bytes", len(file))

	reportLength := len(file)

	payload := map[string]int{"reportLength": reportLength}
	respondWithJSON(w, http.StatusOK, payload)
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
	router.GET("/cpu/:filename", getReportCPU)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running API server in prod mode")
	} else {
		log.Println("Running API server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}
