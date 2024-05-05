package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// making the structs which would act as the database in this project
type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func 

func main() {
	

	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "01", ISBN: "48722", Title: "Heeramandi", Director: &Director{firstname: "Sanjay", lastname: "Bhansali"}})
	movies = append(movies, Movie{ID: "02", ISBN: "58772", Title: "Lapata Ladies", Director: &Director{firstname: "Sanjay", lastname: "Manjrekar"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies[id]", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies[id]", updateMovie).Methods("PUT")
	r.HandleFunc("/movies[id]", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
