package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shared/common/dto"
	"shared/common/utils"

	database_client "client/database"
	db_client "client/database"
	"strconv"

	"strings"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secret-key") // TODO: move to config

func generateJWT(userID int32, role string, deviceID string) (string, error) {
	claims := dto.JWTClaims{
		UserID:   userID,
		Role:     role,
		DeviceID: deviceID,
	}
	// Standard claims
	stdClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	// Merge custom and standard claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, struct {
		dto.JWTClaims
		jwt.RegisteredClaims
	}{
		JWTClaims:        claims,
		RegisteredClaims: stdClaims,
	})
	return token.SignedString(jwtSecret)
}

func parseJWT(tokenStr string) (*dto.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &struct {
		dto.JWTClaims
		jwt.RegisteredClaims
	}{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	if claims, ok := token.Claims.(*struct {
		dto.JWTClaims
		jwt.RegisteredClaims
	}); ok {
		return &claims.JWTClaims, nil
	}
	return nil, jwt.ErrTokenMalformed
}

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

// RegisterAuthUserHandler handles registration of new auth users (admin only)
func RegisterAuthUserHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if req.Role != "admin" && req.Role != "customer" {
		http.Error(w, "Invalid role", http.StatusBadRequest)
		return
	}
	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	// Prepare DB request
	req.Password = string(hash)
	resp, err := database_client.Client.RegisterAuthUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}

// LoginAuthUserHandler handles login for auth users
func LoginAuthUserHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	user, err := database_client.Client.LoginAuthUserHandler(req)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	// Fetch device_id from user record (user.DeviceID)
	deviceID := ""
	if user.DeviceID != "" {
		deviceID = user.DeviceID
	}
	token, err := generateJWT(user.ID, user.Role, deviceID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dto.LoginResponse{Token: token, Role: user.Role})
}

// JWTAuthMiddleware checks for a valid JWT and sets claims in context
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := parseJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AdminOnly middleware allows only admin role
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)
		if !ok || role != "admin" {
			http.Error(w, "Forbidden: admin only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CustomerOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)
		if !ok || role != "customer" {
			http.Error(w, "Forbidden: customer only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
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
