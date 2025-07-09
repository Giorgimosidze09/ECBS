package dto

import (
	"time"
)

type CreateUsersInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserOutput struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Deleted bool   `json:"deleted"`
}

type CountOutput struct {
	Count int `json:"count"`
}

type TotalBalanceOutput struct {
	Total float64 `json:"total"`
}

type UsersListOutput struct {
	ID           int32   `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	CardCount    int64   `json:"card_count"`
	TotalBalance float64 `json:"total_balance"`
	Total        int     `json:"total"`
}

type UsersListInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Card struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CardNumber string `json:"card_number"`
	Active     bool   `json:"active"`
	Type       string `json:"type"`
}

type CardActivation struct {
	ID              int       `json:"id"`
	CardID          int       `json:"card_id"`
	ActivationStart time.Time `json:"activation_start"`
	ActivationEnd   time.Time `json:"activation_end"`
}

type CardOutput struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CardID     string `json:"card_id"`
	DeviceID   int    `json:"device_id"`
	Active     bool   `json:"active"`
	AssignedAt string `json:"assigned_at"`
	Total      int    `json:"total"`
	Deleted    bool   `json:"deleted"`
	Type       string `json:"type"`
}
