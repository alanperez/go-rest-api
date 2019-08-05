package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", home).Methods("GET")
	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: "Hello World!",
	})
}
