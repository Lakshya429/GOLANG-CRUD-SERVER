package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Director *Director `json:"director"`
	Isbn string `json:"isbn"`
}

type Director struct {
	FristName string `json:"FristName"`
	LastName string `json:"LastName"`
}

func getMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)

	for index , item := range movies {
		if item.ID == params["id"]{
		movies  = append(movies[:index] , movies[index+1:]...)
		break
		}
	}
}

func getMovie(w http.ResponseWriter , r  *http.Request){
	w.Header().Set("Content-Type" , "application/json")

	params := mux.Vars(r);

	for  _ , item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies , movie)
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter , r  *http.Request){
	w.Header().Set("content Type" , "application/json")

	params := mux.Vars(r)
	var movie Movie
_ = json.NewDecoder(r.Body).Decode(&movie)
	for _ , item := range movies {
		if item.ID == params["id"] {
			item.Director = movie.Director
			item.Title = movie.Title
			item.Isbn = movie.Isbn
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

var movies []Movie
func main(){
	r := mux.NewRouter()
	
	movies = append(movies , Movie{ID : "1" , Title : "Lengdary Lakshya From Dhapara" , Director : &Director{FristName: "sawam" , LastName: "hum" } , Isbn : "123" } ) ;

	movies = append(movies , Movie{ID : "2" , Title : "Lengdary Lakshya From Dhapara part 2" , Director : &Director{FristName: "sawam" , LastName: "hum" } });
	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMovie).Methods("GET")
	r.HandleFunc("/movies" , createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}" , updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}" , deleteMovie).Methods("DELETE")

	fmt.Println("Server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080" , r))
}