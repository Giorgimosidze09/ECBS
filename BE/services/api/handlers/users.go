package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"shared/common/dto"
	"shared/common/utils"
	"strconv"

	db_client "client/database"

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

func CustomerSumBalanceHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.CustomerSumBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	user, err := db_client.Client.SumBalance(input)
	if err != nil {
		log.Printf("❌ Failed to fetch balance: %v", err)
		http.Error(w, "Failed to fetch balance", http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusCreated, user)
}
