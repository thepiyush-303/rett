package main

import(
	"net/http"
	"fmt"
	"github.com/thepiyush-303/rett/db"

)

func main() {

	dbconn := db.ConnectDB()
	defer dbconn.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/register", db.RegisterUser)

}


func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello world")
}