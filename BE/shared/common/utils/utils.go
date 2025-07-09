package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nats-io/nats.go"
	"golang.org/x/crypto/bcrypt"
)

func Encode[T any](data T) ([]byte, error) {
	return json.Marshal(data)
}

func Decode[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}

func SafeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func SafeFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func SafeBoolean(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}

func SafeInt32(i *int32) int32 {
	if i == nil {
		return 0
	}

	return *i
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func NewUUID() pgtype.UUID {
	return pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}
}

func NullUUID() pgtype.UUID {
	return pgtype.UUID{
		Valid: false,
	}
}

func SafePgUuid(uuid *pgtype.UUID) pgtype.UUID {
	if uuid == nil {
		return pgtype.UUID{}
	}

	return *uuid
}

func ConvertUUID(id *pgtype.UUID) pgtype.UUID {
	if id == nil {
		return pgtype.UUID{Valid: false} // Set Valid to false when nil
	}
	return *id
}

func RespondError(msg *nats.Msg, err error) {
	errorMessage := fmt.Sprintf("Error: %v", err)
	log.Println(errorMessage)
	_ = msg.Respond([]byte(errorMessage))
}

func CheckNATSError(data []byte) error {
	if len(data) >= 5 && string(data[:5]) == "Error" {
		return fmt.Errorf("%s", string(data))
	}
	return nil
}

func ConvertToTimestampTz(str string) pgtype.Timestamptz {
	if strings.TrimSpace(str) == "" {
		return pgtype.Timestamptz{Valid: false}
	}

	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		log.Printf("Invalid timestamp format: %v", err)
		return pgtype.Timestamptz{Valid: false}
	}

	return pgtype.Timestamptz{Time: t, Valid: true}
}

func ValidatePassword(plain string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

func RoundFloatToPrecision(value float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(value*factor) / factor
}

func SafeTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func CommentToJSON[T any](comment T) (string, error) {
	data, err := Encode(comment)
	if err != nil {
		log.Printf("failed to encode comment: %v", err)
		return "", err
	}
	return string(data), nil
}

func StringToPgUUID(s string) (pgtype.UUID, error) {
	var u pgtype.UUID
	err := u.Scan(s)
	if err != nil {
		return pgtype.UUID{Valid: false}, fmt.Errorf("invalid UUID: %w", err)
	}
	return u, nil
}

func IsValidTronAddress(address string) bool {
	if !strings.HasPrefix(address, "T") {
		return false
	}

	decoded := base58.Decode(address)
	return len(decoded) == 21
}

func GetCompanySubject(companyID pgtype.UUID, topic string) string {
	return fmt.Sprintf("company.%s.%s", companyID.String(), topic)
}

func GetDealSubject(companyID pgtype.UUID) string {
	return GetCompanySubject(companyID, "changed.deal")
}

func GetDailyRatesSubject(companyID pgtype.UUID) string {
	return GetCompanySubject(companyID, "daily.rate.refetched")
}

func TimeToPgTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// ParseIDFromRequest extracts an integer ID from the URL path or query (?id=) for REST handlers.
func ParseIDFromRequest(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		idStr = r.URL.Query().Get("id")
	}
	if idStr == "" {
		return 0, fmt.Errorf("missing id parameter")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid id parameter")
	}
	return id, nil
}
