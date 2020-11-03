package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/priyajithc/cart-svc/data"
)

type Cart struct {
	l *log.Logger
}

func CartHandler() *Cart {
	l := log.New(os.Stdout, "cart-svc ", log.LstdFlags)
	return &Cart{l}
}

func (c *Cart) AddToCart(w http.ResponseWriter, r *http.Request) {
	// fetch the products from the datastore

	c.l.Println("Inside AddToCart")

	idInStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idInStr)
	c.l.Println("Id", id)

	var cps []data.CartProduct
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Data not found", http.StatusBadRequest)
	}

	json.Unmarshal(reqBody, &cps)
	errAddToCart := data.AddToCart(id, cps)
	if errAddToCart != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cps)
	}
}

func (c *Cart) GetCart(w http.ResponseWriter, r *http.Request) {
	// fetch the products from the datastore

	c.l.Println("Inside GetCart")

	idInStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idInStr)
	c.l.Println("Id", id)

	ct, err := data.GetCart(id)
	if err != nil {
		http.Error(w, "Unable to find Cart", http.StatusNotFound)
	} else {
		// serialize the list to JSON
		err := ct.ToJSON(w)
		if err != nil {
			http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		}
	}

}
