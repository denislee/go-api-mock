package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fileName := strings.Replace(r.RequestURI, "/", "_", -1)
	b, err := ioutil.ReadFile("responses/" + fileName + ".txt")
	fileContent := string(b)
	if err != nil {
		fileContent = err.Error()
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fileContent))
}

func main() {
	http.HandleFunc("/", apiResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
