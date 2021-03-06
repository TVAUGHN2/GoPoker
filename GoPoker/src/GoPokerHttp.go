package main

import(
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/handvalue", handValue).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", router))
}