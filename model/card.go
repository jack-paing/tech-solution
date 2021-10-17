package model

type Card struct {
	CardID       string   `json:"card_id" db:"card_id"`
	Name         *string  `json:"name" db:"name"`
	Desc         *string  `json:"desc" db:"description"`
	DailyLimit   *float64 `json:"daily_limit" db:"daily_limit"`
	MonthlyLimit *float64 `json:"monthly_limit" db:"monthly_limit"`
}
