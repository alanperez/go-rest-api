package controllers

import (
	"net/http"

	"github.com/alanperez/go-rest-api/utils"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	utils.ToJson(w, struct {
		Message string `json: "message"`
	}{
		Message: "Go RESTful API",
	})
}
