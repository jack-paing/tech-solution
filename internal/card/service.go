package card

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"tech-solution/internal/repo"
	"tech-solution/model"
)

type Service interface {
	Create(ctx context.Context, card *model.Card) error
}

type service struct {
	repo *repo.Card
}

func NewService(db *sqlx.DB) Service {
	return &service{
		repo: repo.NewCard(db),
	}
}

func (s *service) Create(ctx context.Context, card *model.Card) error {
	cardUID := uuid.New().String()
	card.CardID = cardUID
	return s.repo.Create(ctx, card)
}
