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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ReadResponse, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}



func main(){
	r := mux.NewRouter

	movies = append(movies, Movie{ID: "1", Isbn:"438227", Title: "Movie One", Director: &Director{Firstname:"Hamid", Lastname: "Reazei"} })
	movies = append(movies, Movie{ID: "2", Isbn:"45455", Title: "Movie Two", Director: &Director{Firstname:"Amir Bahador", Lastname: "Bahadori"} })
	r.HandleFunc("./movies" , getMovie).Methods("GET")
	r.HandleFunc("./movies/{id}" , getMovie).Methods("GET")
	r.HandleFunc("./movies" , createMovie).Methods("POST")
	r.HandleFunc("./movies/{id}" , updateMovie).Methods("PUT")
	r.HandleFunc("./movies/{id}" , deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}