package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := fmt.Sprintf("Hello, %s!", vars["PARAM"])
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, response)
}

func bad(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func data(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Println(body)
	response := fmt.Sprintf("I got message:\n%s", body)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, response)

}

func headers(w http.ResponseWriter, r *http.Request) {
	firstVal, _ := strconv.Atoi(r.Header.Get("a"))
	secondVal, _ := strconv.Atoi(r.Header.Get("b"))

	resSum := firstVal + secondVal
	
	w.Header().Set("a+b", fmt.Sprintf("%d", resSum))
}

func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", hello).Methods("GET")
	router.HandleFunc("/bad", bad).Methods("GET")
	router.HandleFunc("/data", data).Methods("POST")
	router.HandleFunc("/headers", headers).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	fmt.Println(host)
	port, err := strconv.Atoi(os.Getenv("PORT"))
	fmt.Println(port)
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
