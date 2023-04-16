package server

import (
	"encoding/json"
	"net/http"

	"github.com/worthlessowl/erajaya/usecase"
)

func (server *Server) HandleInsertProduct(w http.ResponseWriter, r *http.Request) {
	product := usecase.AddProduct{}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "cant decode body", http.StatusBadRequest)
		return
	}

	err = server.uc.InsertProduct(product)
	if err != nil {
		http.Error(w, "something went wrong in our server", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("success"))
}

func (server *Server) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	sortByQuery := r.URL.Query().Get("sortby")
	sortBy := -1

	if sortByQuery == "newest" {
		sortBy = int(usecase.Newest)
	}
	if sortByQuery == "cheapest" {
		sortBy = int(usecase.Cheapest)
	}
	if sortByQuery == "mostexpensive" {
		sortBy = int(usecase.MostExpensive)
	}
	if sortByQuery == "nameascending" {
		sortBy = int(usecase.NameAsc)
	}
	if sortByQuery == "namedescending" {
		sortBy = int(usecase.NameDesc)
	}

	products, err := server.uc.ListProduct(usecase.Filter(sortBy))
	if err != nil {
		http.Error(w, "something went wrong in our server", http.StatusInternalServerError)
		return
	}

	marshalled, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "something went wrong in our server", http.StatusInternalServerError)
		return
	}

	w.Write(marshalled)
}
