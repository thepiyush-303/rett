package main

import(
	"net/http"
	"fmt"
	"github.com/thepiyush-303"
)

func main() {

	db := 

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/register", handleRegister)

}

func handleRegister(w http.ResponseWriter, r *http.Request){

}

func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello world")
}