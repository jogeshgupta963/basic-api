package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func seedData(){
	movies = append(movies,Movie{ID:"1",Isbn:"askmd",Title:"Movie-1",Director:&Director{Firstname: "John",Lastname: "Doe"}});
	movies = append(movies,Movie{ ID:"2", Isbn:"zxcvbnm",Title:"Movie-2",Director:&Director{Firstname: "Jane",Lastname: "Doe"}});

}

func getAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies);
	
}

func deleteMovie(res http.ResponseWriter, req *http.Request){
	params := mux.Vars(req);
	
	for index,item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break;
		}
	}

	res.Header().Set("Content-Type","application/json")
}

func getMovie (res http.ResponseWriter,req *http.Request){

	params := mux.Vars(req);
	var movie Movie
	for _,item := range movies {
		if item.ID == params["id"]{
			movie = item
			break;
		}
	}
	res.Header().Set("Content-Type","application/json")
	json.NewEncoder(res).Encode(movie);
}
func createMovie(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","application/json")
	var x Movie;
	if req.Body == nil {
		json.NewEncoder(res).Encode("Invalid data")
		return
	}
	_ = json.NewDecoder(req.Body).Decode(&x);
	movies = append(movies, x)
	json.NewEncoder(res).Encode(x);
}
func main() {

	seedData()
	r := mux.NewRouter()
	r.HandleFunc("/movies/get-all",getAllMovies).Methods("GET")
	r.HandleFunc("/movies/create/",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	
	fmt.Println("Server Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000",r))

}