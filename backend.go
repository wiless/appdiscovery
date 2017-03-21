package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
)

type App struct {
	AppName     string
	Source      string
	Description string
}

var registeredApps []App

func init() {
	registeredApps = make([]App, 5)
	for i := 0; i < 5; i++ {
		registeredApps[i] = App{"Hello #" + strconv.FormatInt(int64(i), 10), "github.com/wiless/cellular", "Its a blah blah App"}
	}
}

func main() {
	log.Print("Started...")
	log.Print("Visit http://localhost:8080")
	router := mux.NewRouter().StrictSlash(true)

	// Set API paths
	router.HandleFunc("/api/apps", handleApps).Methods("GET")
	router.HandleFunc("/api/apps/{appid:[0-9]+}", getApp).Methods("GET")

	//Set Statick service
	//Approach 1
	router.PathPrefix("/www").HandlerFunc(Index)
	//Approah 2
	//router.PathPrefix("/www").Handler(http.StripPrefix("/www", http.FileServer(http.Dir("."))))

	// Set API paths
	router.HandleFunc("/", welcomePage)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getApp(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	log.Print("looking for an App # ", vars)
	val, ok := vars["appid"]
	_ = val
	if !ok {
		fmt.Fprintf(w, "Dont know what are you asking")
	}

	b, er := json.Marshal(registeredApps[cast.ToInt(val)])
	if er != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set(key string, value string)
	fmt.Fprintf(w, "%s", b)

}

func handleApps(w http.ResponseWriter, r *http.Request) {
	log.Printf("Let me tell you about my registered Apps \n %#v", r.Context())

	b, er := json.Marshal(registeredApps)
	if er != nil {
		http.Error(w, "Error converting results to json", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set(key string, value string)
	fmt.Fprintf(w, "%s", b)

}
