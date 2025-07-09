package dto

type TopUpInput struct {
	UserID   int     `json:"user_id"`
	CardID   int     `json:"card_id"`
	Balance  float64 `json:"balance"`
	RideCost float64 `json:"ride_cost"`
}

type BalanceOutput struct {
	UserID    int     `json:"user_id"`
	CardID    int     `json:"card_id"`
	Balance   float64 `json:"balance"`
	RideCost  float64 `json:"ride_cost"`
	UpdatedAt string  `json:"updated_at"`
}
