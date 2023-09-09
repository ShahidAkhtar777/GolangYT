package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Learning about MOD files in Golang!!")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey go mod users!!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang series on YT.</h1>"))
}

// We dont usually touch the mod file 
// To edit use:
// go mod edit -go 1.16
// go mod edit -module <newname>
// running go mod vendor creates a vendor folder same as a node modules file
// Now our dependencies gets used from that folder and not from the cache
// Recommendation is not use vendor just use the dependencies from the go.mod file
