package models


type Bet struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	SportID  int     `json:"sport_id"`
	Amount   float64 `json:"amount"`
	CreateAt string  `json:"created_at"`
}