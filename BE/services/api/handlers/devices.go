package handlers

import (
	database "client/database"

	"encoding/json"
	"log"
	"net/http"
	"shared/common/dto"
	"shared/common/utils"
)

func CreateDevices(w http.ResponseWriter, r *http.Request) {
	var input dto.DevicesInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedBalance, err := database.Client.CreateDevices(input)
	if err != nil {
		log.Printf("Failed to create devices: %v", err)
		http.Error(w, "Failed to create devices", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, updatedBalance)
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
	total, err := database.Client.DevicesList(input)
	if err != nil {
		log.Printf("Failed to get devices list: %v", err)
		http.Error(w, "Failed to get devices list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(total)
}

func GetDeviceByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}
	device, err := database.Client.GetDeviceByID(id)
	if err != nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}
	utils.RespondJSON(w, http.StatusOK, device)
}

func UpdateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var input dto.DeviceOutput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	device, err := database.Client.UpdateDevice(input)
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
	err = database.Client.SoftDeleteDevice(id)
	if err != nil {
		http.Error(w, "Failed to delete device", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
