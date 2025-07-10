package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"shared/common/dto"

	database_client "client/database"

	"strings"

	"time"

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
