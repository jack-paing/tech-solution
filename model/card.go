package model

import "time"

type Card struct {
	CardID       string    `json:"card_id" db:"card_id"`
	Name         *string   `json:"name" db:"name"`
	Desc         *string   `json:"description" db:"description"`
	DailyLimit   *float64  `json:"daily_limit" db:"daily_limit"`
	MonthlyLimit *float64  `json:"monthly_limit" db:"monthly_limit"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
