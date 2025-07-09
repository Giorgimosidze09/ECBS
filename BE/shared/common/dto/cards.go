package dto

type AssignCardInput struct {
	UserID          int    `json:"user_id"`
	CardID          string `json:"card_id"`
	DeviceID        int    `json:"device_id"`
	Type            string `json:"type"` // 'balance' or 'activation'
	ActivationStart string `json:"activation_start,omitempty"`
	ActivationEnd   string `json:"activation_end,omitempty"`
}

type ValidateCardInput struct {
	CardID int `json:"card_id"`
}

type RideCostInput struct {
	RideCost float64 `json:"ride_cost"`
}
type ValidateCardOutput struct {
	Valid    bool    `json:"valid"`
	UserID   int     `json:"user_id,omitempty"`
	UserName string  `json:"user_name,omitempty"`
	Balance  float64 `json:"balance,omitempty"`
	Message  string  `json:"message,omitempty"`
}

type Charges struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	Total       int     `json:"total"`
}
