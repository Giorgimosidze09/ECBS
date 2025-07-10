package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shared/common/dto"
	"shared/common/utils"

	database_client "client/database"
	db_client "client/database"
)

func SyncAuthorizedAccessHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.AuthorizedInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db_client.Client.SyncAuthorizedAccess(input)
	if err != nil {
		log.Printf("Failed to validate card: %v", err)
		http.Error(w, "Failed to validate card", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, result)
}

func SyncAccessLogs(w http.ResponseWriter, r *http.Request) {
	var input dto.SyncAccessLogInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if len(input.Logs) == 0 {
		http.Error(w, "No logs to sync", http.StatusBadRequest)
		return
	}

	if err := database_client.Client.SyncAccessLogs(input); err != nil {
		http.Error(w, fmt.Sprintf("Sync failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"synced"}`))
}
