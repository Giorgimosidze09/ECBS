package handlers

import (
	db_client "client/database"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shared/common/dto"
	"shared/common/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// SMSProvider defines the interface for sending SMS
// You can implement this for Twilio, local providers, etc.
type SMSProvider interface {
	Send(to string, message string) error
}

// DefaultSMSProvider just logs the SMS (for development)
type DefaultSMSProvider struct{}

func (d DefaultSMSProvider) Send(to string, message string) error {
	fmt.Printf("[SMS] To: %s | Message: %s\n", to, message)
	return nil
}

// Global SMS provider instance (swap this for a real provider)
var smsProvider SMSProvider = DefaultSMSProvider{}

// SendSMS uses the global provider
func SendSMS(to string, message string) error {
	// TODO: Replace smsProvider with a real implementation for your country/provider
	return smsProvider.Send(to, message)
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

	// Fetch user and send SMS
	user, err := db_client.Client.GetUserByID(createdCard.UserID)
	if err == nil && user != nil && user.Phone != "" {
		msg := fmt.Sprintf("Your card %s has been assigned. Type: %s.", createdCard.CardID, createdCard.Type)
		SendSMS(user.Phone, msg)
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

	// Fetch user and send SMS
	user, err := db_client.Client.GetUserByID(updatedBalance.UserID)
	if err == nil && user != nil && user.Phone != "" {
		msg := fmt.Sprintf("Your card %d has been topped up. New balance: %.2f.", updatedBalance.CardID, updatedBalance.Balance)
		SendSMS(user.Phone, msg)
	}

	utils.RespondJSON(w, http.StatusCreated, updatedBalance)
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

func PayboxTopupHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.PayboxTopupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if req.CardID <= 0 || req.Amount <= 0 {
		http.Error(w, "Missing card_id or invalid amount", http.StatusBadRequest)
		return
	}
	err := db_client.Client.AddBalanceToCard(req)
	if err != nil {
		http.Error(w, "Failed to update balance", http.StatusInternalServerError)
		return
	}
	// Fetch card to get user_id, then fetch user and send SMS
	card, err := db_client.Client.GetCardByID(req.CardID)
	if err == nil && card != nil {
		user, err := db_client.Client.GetUserByID(card.UserID)
		if err == nil && user != nil && user.Phone != "" {
			msg := fmt.Sprintf("Your card %s has been topped up via PayBox. Amount: %.2f.", card.CardID, req.Amount)
			SendSMS(user.Phone, msg)
		}
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Top-up successful",
		"card_id": req.CardID,
		"amount":  req.Amount,
	})
}
