package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shared/common/dto"

	db_client "client/database"
)

func GetUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db_client.Client.CountUsers()
	if err != nil {
		log.Printf("Failed to get user count: %v", err)
		http.Error(w, "Failed to get user count", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(count)
}

func GetCardStatsHandler(w http.ResponseWriter, r *http.Request) {
	count, err := db_client.Client.CountCards()
	if err != nil {
		log.Printf("Failed to get card count: %v", err)
		http.Error(w, "Failed to get card count", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(count)
}

func GetTotalBalanceHandler(w http.ResponseWriter, r *http.Request) {
	total, err := db_client.Client.TotalBalance()
	if err != nil {
		log.Printf("Failed to get total balance: %v", err)
		http.Error(w, "Failed to get total balance", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(total)
}

func GetCharges(w http.ResponseWriter, r *http.Request) {
	var input dto.UsersListInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if input.Limit <= 0 {
		input.Limit = 25 // Default limit
	}
	if input.Offset < 0 {
		input.Offset = 0 // Default offset
	}
	total, err := db_client.Client.GetCharges(input)
	if err != nil {
		log.Printf("Failed to get Charges list: %v", err)
		http.Error(w, "Failed to get Charges list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(total)
}
