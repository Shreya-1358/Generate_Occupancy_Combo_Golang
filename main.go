package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project_3/model"
)

func main() {
	fmt.Println("Listening at server 8080:")
	http.HandleFunc("/generate_occupancy", generateOccupancy)
	http.ListenAndServe(":8080", nil)

}

func generateOccupancy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.OccupancyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	occup := model.NewOccupancyImpl()
	resp := occup.GenerateOccupancy(&req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
