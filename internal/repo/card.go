package repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"tech-solution/model"
	"time"
)

const cardAttrs = "card_id,name,description,daily_limit,monthly_limit,created_at,updated_at"

type CardRepo interface {
	Create(ctx context.Context, card *model.Card) error
	Get(ctx context.Context) ([]*model.Card, error)
	GetByCardId(ctx context.Context, cardId string) (*model.Card, error)
	Delete(ctx context.Context, cardId string) error
	Update(ctx context.Context, cardId string, updateCard *model.Card) error
}

type Card struct {
	db *sqlx.DB
}

func NewCard(db *sqlx.DB) *Card {
	return &Card{db: db}
}

func (r Card) Create(ctx context.Context, card *model.Card) error {
	query := `
		INSERT INTO cards 
			(card_id,name,description,daily_limit,monthly_limit)
		VALUES
			(?,?,?,?,?)`
	_, err := r.db.ExecContext(ctx, query, card.CardID, card.Name, card.Desc, card.DailyLimit, card.MonthlyLimit)
	return err
}

func (r Card) Get(ctx context.Context) ([]*model.Card, error) {
	query := fmt.Sprintf(`
	SELECT %s FROM cards;
	`, cardAttrs)
	var cards []*model.Card
	if err := r.db.SelectContext(ctx, &cards, query); err != nil {
		return nil, err
	}
	return cards, nil
}

func (r Card) GetByCardId(ctx context.Context, cardId string) (*model.Card, error) {
	query := fmt.Sprintf(`
	SELECT %s FROM cards WHERE card_id = '%s';
	`, cardAttrs, cardId)
	card := &model.Card{}
	if err := r.db.GetContext(ctx, card, query); err != nil {
		return nil, err
	}
	return card, nil
}

func (r Card) Delete(ctx context.Context, cardId string) error {
	query := fmt.Sprintf(`
		DELETE FROM cards WHERE card_id = '%s';
		`, cardId)
	_, err := r.db.ExecContext(ctx, query)
	return err
}

func (r Card) Update(ctx context.Context, cardId string, updateCard *model.Card) error {
	query := fmt.Sprintf(`
		UPDATE cards
		SET
		name=COALESCE(?,name),
		description=COALESCE(?,description),
		daily_limit=COALESCE(?,daily_limit),
		monthly_limit=COALESCE(?,monthly_limit),
		updated_at=?
		WHERE card_id='%s'
		`, cardId)
	_, err := r.db.ExecContext(ctx, query, updateCard.Name, updateCard.Desc, updateCard.DailyLimit, updateCard.MonthlyLimit, time.Now())
	return err
}
