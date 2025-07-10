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
	"strconv"

	"github.com/go-chi/chi/v5"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUsersInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := db_client.Client.CreateUser(input)
	if err != nil {
		log.Printf("❌ Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, user)
}

func AssignCardHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.AssignCardInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdCard, err := db_client.Client.CreateCard(input)
	if err != nil {
		log.Printf("Failed to assign card: %v", err)
		http.Error(w, "Failed to assign card", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdCard)
}

func TopUpBalanceHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.TopUpInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedBalance, err := db_client.Client.TopUpBalance(input)
	if err != nil {
		log.Printf("Failed to top up balance: %v", err)
		http.Error(w, "Failed to top up balance", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, updatedBalance)
}

func CreateDevices(w http.ResponseWriter, r *http.Request) {
	var input dto.DevicesInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedBalance, err := db_client.Client.CreateDevices(input)
	if err != nil {
		log.Printf("Failed to create devices: %v", err)
		http.Error(w, "Failed to create devices", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, updatedBalance)
}

func ValidateCardHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.ValidateCardInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db_client.Client.ValidateCard(input)
	if err != nil {
		log.Printf("Failed to validate card: %v", err)
		http.Error(w, "Failed to validate card", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, result)
}

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

func HandleCardScanWebhook(w http.ResponseWriter, r *http.Request) {
	var input dto.ValidateCardInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db_client.Client.ValidateCard(input)
	if err != nil {
		log.Printf("Failed to validate card: %v", err)
		http.Error(w, "Failed to validate card", http.StatusInternalServerError)
		return
	}

	resp := struct {
		Valid bool `json:"valid"`
	}{
		Valid: result.Valid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

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

func GetUserList(w http.ResponseWriter, r *http.Request) {
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

	total, err := db_client.Client.GetUsersList(input)
	if err != nil {
		log.Printf("Failed to get users list: %v", err)
		http.Error(w, "Failed to get users list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(total)
}

func GetCardsList(w http.ResponseWriter, r *http.Request) {
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

	total, err := db_client.Client.GetCardsList(input)
	if err != nil {
		log.Printf("Failed to get cards list: %v", err)
		http.Error(w, "Failed to get cards list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(total)
}

func BalanceList(w http.ResponseWriter, r *http.Request) {
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

	total, err := db_client.Client.BalanceList(input)
	if err != nil {
		log.Printf("Failed to get balance list: %v", err)
		http.Error(w, "Failed to get balance list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(total)
}

func DevicesList(w http.ResponseWriter, r *http.Request) {
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

	total, err := db_client.Client.DevicesList(input)
	if err != nil {
		log.Printf("Failed to get devices list: %v", err)
		http.Error(w, "Failed to get devices list", http.StatusInternalServerError)
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

func ChangeRideCost(w http.ResponseWriter, r *http.Request) {
	var input dto.RideCostInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := db_client.Client.ChangeRideCost(input)
	if err != nil {
		log.Printf("Failed to change ride cost: %v", err)
		http.Error(w, "Failed to change ride cost", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, result)
}

func AddCardActivationHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.CardActivation
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	result, err := db_client.Client.AddCardActivation(input)
	if err != nil {
		log.Printf("Failed to add card activation: %v", err)
		http.Error(w, "Failed to add card activation", http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusCreated, result)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.UserOutput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	idStr := chi.URLParam(r, "id")
	if idStr != "" {
		if id, err := strconv.Atoi(idStr); err == nil {
			input.ID = int32(id)
		}
	}
	user, err := db_client.Client.UpdateUser(input)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, user)
}

func SoftDeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	err = db_client.Client.SoftDeleteUser(id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := db_client.Client.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	utils.RespondJSON(w, http.StatusOK, user)
}

func UpdateCardHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.CardOutput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	idStr := chi.URLParam(r, "id")
	if idStr != "" {
		if id, err := strconv.Atoi(idStr); err == nil {
			input.ID = id
		}
	}
	card, err := db_client.Client.UpdateCard(input)
	if err != nil {
		http.Error(w, "Failed to update card", http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, card)
}

func SoftDeleteCardHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid card ID", http.StatusBadRequest)
		return
	}
	err = db_client.Client.SoftDeleteCard(id)
	if err != nil {
		http.Error(w, "Failed to delete card", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetCardByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid card ID", http.StatusBadRequest)
		return
	}
	card, err := db_client.Client.GetCardByID(id)
	if err != nil {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}
	utils.RespondJSON(w, http.StatusOK, card)
}

func UpdateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.DeviceOutput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	device, err := db_client.Client.UpdateDevice(input)
	if err != nil {
		http.Error(w, "Failed to update device", http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusOK, device)
}

func SoftDeleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}
	err = db_client.Client.SoftDeleteDevice(id)
	if err != nil {
		http.Error(w, "Failed to delete device", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetDeviceByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}
	device, err := db_client.Client.GetDeviceByID(id)
	if err != nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}
	utils.RespondJSON(w, http.StatusOK, device)
}
