package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"frontendmasters.com/go/musuem/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"]
	w.Header().Set("Header", "application/json")

	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.Getall()) {
			json.NewEncoder(w).Encode(data.Getall()[finalId])

		} else  {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.Getall())
	}
}