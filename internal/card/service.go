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
	Get(ctx context.Context) ([]*model.Card, error)
	GetByCardId(ctx context.Context, cardId string) (*model.Card, error)
	Delete(ctx context.Context, cardId string) error
	Update(ctx context.Context, cardId string, card *model.Card) error
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

func (s *service) Get(ctx context.Context) ([]*model.Card, error) {
	return s.repo.Get(ctx)
}

func (s *service) GetByCardId(ctx context.Context, cardId string) (*model.Card, error) {
	return s.repo.GetByCardId(ctx, cardId)
}

func (s *service) Delete(ctx context.Context, cardId string) error {
	return s.repo.Delete(ctx, cardId)
}

func (s *service) Update(ctx context.Context, cardId string, card *model.Card) error {
	return s.repo.Update(ctx, cardId, card)
}
