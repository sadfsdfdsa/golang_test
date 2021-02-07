package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("id=%s", id)
	fmt.Fprint(w, response)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Page")
}

func main() {
	port := os.Getenv("PORT")
    // if port == nil {
    //     port = "8181"
    // }

	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", productsHandler).Methods("GET")
	router.HandleFunc("/articles/{id:[0-9]+}", productsHandler).Methods("POST")
	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+port, nil)
}
