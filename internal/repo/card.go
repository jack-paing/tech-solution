package repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"tech-solution/model"
)

type Card struct {
	db *sqlx.DB
}

func NewCard(db *sqlx.DB) *Card {
	return &Card{db: db}
}

func (r Card) Create(ctx context.Context, card *model.Card) error {
	fmt.Println(r.db)
	query := `
		INSERT INTO cards 
			(card_id,name,description,daily_limit,monthly_limit)
		VALUES
			(?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, card.CardID, card.Name, card.Desc, card.DailyLimit, card.MonthlyLimit)
	fmt.Println(err)
	return err
}
