package main

import (
	"encoding/json" // needed to encode data to json when we send it to postman
	"fmt"           //for printing out stuff like your server is connected and all
	"log"           // for log out the errors if there's any error for connecting to the server
	"math/rand"     //when creating a new id or something like that
	"net/http"      //create server in golang
	"strconv"       // math.rand will crreate an integer and "strconv" will convert that into string

	"github.com/gorilla/mux"
)

// we are using struct and slices in this program
// it is like an object in JS or JAVA or in any other language.
//we're going to have key-value pairs and define the types of data inside that

//Movie and Director are going to be associated

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // star is a pointer that means if i create a struct called director it
	//  will be associated  to movie struct

}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Now we have defined two structs Movie and Director and now we'll define a variable called movies which will be slice of the type movie

var movies []Movie

func main() {
	//NewRouter is a function inside the mux library inside gorilla
	r := mux.NewRouter() // := it declare and define the variable in the same time in golang so now r is our NewRouter

	movies = append(movies, Movie{ID: "1", ISBN: "453654", Title: "Ra-One", Director: &Director{Firstname: "Nittyansh", Lastname: "Srivastava"}})
	movies = append(movies, Movie{ID: "2", ISBN: "456723", Title: "MYLOVE", Director: &Director{Firstname: "Shivansh", Lastname: "Srivastava"}})
	//we want the reference of the director that's why Director: &Director is used becaue here * is a pointer
	//Now we have 2 movies so if we check our code then we can check our API is working fine or not.

	// & is used tot give you the address and * is used to acess the address

	r.HandleFunc("/movies", getMovies).Methods("GET") //using get method for getting all the movies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	//now to start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}

//NOW we have the function, methods and the routes

// Let's create getMovies() function ----------->>>>              (1)
func getMovies(w http.ResponseWriter, r *http.Request) {

	// passing a pointer of the request that i will send from my postman to this function and w is the response writer

	w.Header().Set("Content-Type", "application/json") //basically we want to set the content type as json
	json.NewEncoder(w).Encode(movies)                  // we're going to encode w basically the response that we want to send it and encode it into json
}

// let's create the delete function ----------->>>>>		(2)
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) // get some params so that the id that'll pass from postman which will go as a param to this function createMovie and that param which will be the id will be present mux.vars and inside the request which is r so it will part of the request the pointer to the request i'm sending out here so in params now i have the now i'll be able to access the id
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// let's create the getMovie Function ----------->>>>>		(3)
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//creating the CreateMovie function ------------->>> (4)

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) //telling user that we have created this movie
}

//Updating movie functon --------------->>>>>>>>     (5)

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("content-type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, ramge
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
