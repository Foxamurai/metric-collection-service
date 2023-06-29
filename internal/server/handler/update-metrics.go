package handler

import (
	"fmt"
	"log"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println(r.URL)
}

func UpdateMetrics(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log.Print("Update message received")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(nil)
	}

	// TODO
}
