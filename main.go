package main

import(
	"net/http"
	"fmt"	
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/register", handleRegister)

}

func handleRegister(w http.ResponseWriter, r *http.Request){
	
}

func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello world")
}