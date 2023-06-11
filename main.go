package main 

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID 			string `json:"id"`
	Isbn 		string `json:"isbn"`
	Title  		string `json:"title"`
	Director 	*Director `json:"director"`
}

type Director struct{
	Firstname 	string `json:"firstname"`  
	Lastname	string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main(){
	r := mux.NewRouter

	movies = append(movies, Movie{ID: "1", Isbn:"438227", Title: "Movie One", Director: &Director{Fisrtname:"Hamid", lastname: "Reazei"} })
	movies = append(movies, Movie{ID: "2", Isbn:"45455", Title: "Movie Two", Director: &Director{Fisrtname:"Amir Bahador", lastname: "Bahadori"} })
	r.HandleFunc("./movies" , getMovies).Methods("GET")
	r.HandleFunc("./movies/{id}" , getMovies).Methods("GET")
	r.HandleFunc("./movies" , createMovies).Methods("POST")
	r.HandleFunc("./movies/{id}" , updateMovies).Methods("PUT")
	r.HandleFunc("./movies/{id}" , deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}