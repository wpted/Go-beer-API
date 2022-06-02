package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Beer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Brand *Brand `json:"brand"`
}

type Brand struct {
	CompanyName string `json:"company-name"`
	Country     string `json:"country"`
}

var beers []Beer

// getMovies is a UrlHandlerFunction that tells Gorilla to match HTTP GET method
func getBeers(w http.ResponseWriter, r *http.Request) {
	//	http.ResponseWriter is an interface that implements io.writer and used for sending the server response
	w.Header().Set("Content-Type", "application/json")
	// NewEncoder returns a new encoder that writes to w
	json.NewEncoder(w).Encode(beers)
}

func deleteBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range beers {

		if item.ID == params["id"] {
			beers = append(beers[:index], beers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(beers)
}

func getBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range beers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var beer Beer
	// Parsing the user parameters(reading data from r) and decode it
	_ = json.NewDecoder(r.Body).Decode(&beer)
	beer.ID = strconv.Itoa(rand.Intn(10000000))
	beers = append(beers, beer)
	json.NewEncoder(w).Encode(beer)
}
func updateBeer(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	// loop over the beers, range
	for index, item := range beers {
		if item.ID == params["id"] {
			// delete the beer with the i.d that the user sent
			beers = append(beers[:index], beers[index+1:]...)
			// add a new beer
			var beer Beer
			_ = json.NewDecoder(r.Body).Decode(&beer)
			beer.ID = item.ID
			beers = append(beers, beer)
			json.NewEncoder(w).Encode(beer)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	// Hardcode beers to the slice
	beers = append(beers,
		Beer{ID: "1", Name: "Taiwan Beer", Price: "35 NTD", Brand: &Brand{CompanyName: "TaiwanBeer", Country: "Taiwan"}})
	beers = append(beers,
		Beer{ID: "2", Name: "Tiger Beer", Price: "35 NTD", Brand: &Brand{CompanyName: "Tiger", Country: "Singapore"}})

	// we have several URLs that should only match when the host is :8000
	r.HandleFunc("/beers", getBeers).Methods("GET")
	r.HandleFunc("/beers/{id}", getBeer).Methods("GET")
	r.HandleFunc("/beers", createBeer).Methods("POST")
	r.HandleFunc("/beers/{id}", updateBeer).Methods("PUT")
	r.HandleFunc("/beers/{id}", deleteBeer).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
