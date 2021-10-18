package card

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"tech-solution/internal/repo/mocks"
	"tech-solution/model"
	ut "tech-solution/utils"
	"testing"
)

func TestCreateService(t *testing.T) {
	tests := []struct {
		name  string
		card  *model.Card
		error error
	}{
		{"Valid Data", &model.Card{
			Name:         ut.PtrString("Test card"),
			Desc:         ut.PtrString("Test Desc"),
			DailyLimit:   ut.PtrFloat(100.50),
			MonthlyLimit: ut.PtrFloat(3000.90),
		}, nil},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		cardRepoMock := &mocks.CardRepo{}
		cardRepoMock.On("Create", ctx, tt.card).Return(tt.error)

		service := NewService(cardRepoMock)
		t.Run(tt.name, func(t *testing.T) {
			err := service.Create(ctx, tt.card)
			assert.Equal(t, err, tt.error)
		})
	}
}

func TestDeleteService(t *testing.T) {
	tests := []struct {
		name   string
		cardId string
		error  error
	}{
		{"Valid card id", "fd439dc5-024a-4003-952e-56aee4b39255", nil},
		{"Invalid card id", "fd439dc5-024a-4003-952e-56a", errors.New("invalid")},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		cardRepoMock := &mocks.CardRepo{}
		cardRepoMock.On("Delete", ctx, tt.cardId).Return(tt.error)

		service := NewService(cardRepoMock)
		t.Run(tt.name, func(t *testing.T) {
			err := service.Delete(ctx, tt.cardId)
			assert.Equal(t, err, tt.error)
		})
	}
}

func TestGetService(t *testing.T) {
	tests := []struct {
		name       string
		cardResult []*model.Card
		error      error
	}{
		{"Get Cards", make([]*model.Card, 2), nil},
		{"Invalid ", nil, errors.New("invalid")},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		cardRepoMock := &mocks.CardRepo{}
		cardRepoMock.On("Get", ctx).Return(tt.cardResult, tt.error)

		service := NewService(cardRepoMock)
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.Get(ctx)
			assert.Equal(t, err, tt.error)
		})
	}
}

func TestGetByIdService(t *testing.T) {
	tests := []struct {
		name       string
		cardResult *model.Card
		cardId     string
		error      error
	}{
		{"Get Card by id", &model.Card{
			CardID:       "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada",
			Name:         ut.PtrString("card name 2"),
			Desc:         ut.PtrString("Test card"),
			DailyLimit:   ut.PtrFloat(100.50),
			MonthlyLimit: ut.PtrFloat(3000)}, "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada", nil},
		{"Invalid ", nil, "585cbfc0-41dd-4e9f-b7aa-3f32", errors.New("invalid id")},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		cardRepoMock := &mocks.CardRepo{}
		cardRepoMock.On("GetByCardId", ctx, tt.cardId).Return(tt.cardResult, tt.error)

		service := NewService(cardRepoMock)
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.GetByCardId(ctx, tt.cardId)
			assert.Equal(t, err, tt.error)
		})
	}
}

func TestUpdateService(t *testing.T) {
	tests := []struct {
		name          string
		updateCardReq *model.Card
		cardId        string
		error         error
	}{
		{"Get Card by id", &model.Card{
			CardID:       "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada",
			Name:         ut.PtrString("card name 2"),
			Desc:         ut.PtrString("Test card"),
			DailyLimit:   ut.PtrFloat(100.50),
			MonthlyLimit: ut.PtrFloat(3000)}, "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada", nil},
		{"Invalid ", nil, "585cbfc0-41dd-4e9f-b7aa-3f32", errors.New("invalid id")},
	}
	ctx := context.TODO()
	for _, tt := range tests {
		cardRepoMock := &mocks.CardRepo{}
		cardRepoMock.On("Update", ctx, tt.cardId, tt.updateCardReq).Return(tt.error)

		service := NewService(cardRepoMock)
		t.Run(tt.name, func(t *testing.T) {
			err := service.Update(ctx, tt.cardId, tt.updateCardReq)
			assert.Equal(t, err, tt.error)
		})
	}
}
