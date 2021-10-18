package repo

import (
	"context"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"tech-solution/model"
	ut "tech-solution/utils"
	"testing"
)

func setupCardRepoDB(t *testing.T) *sqlx.DB {
	dbURL := "root:yourpasswd@tcp(127.0.0.1:3306)/user_service_test?multiStatements=true&parseTime=true"
	mysqlDB, err := sqlx.Open("mysql", dbURL)
	require.NoError(t, err)

	_, err = mysqlDB.ExecContext(context.Background(), `TRUNCATE TABLE cards`)
	require.NoError(t, err)

	return mysqlDB
}

func TestCreateCard(t *testing.T) {
	tests := []struct {
		name  string
		card  *model.Card
		error error
	}{
		{"Valid card request", &model.Card{CardID: "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada", Name: ut.PtrString("Testing name"), DailyLimit: ut.PtrFloat(100.50), MonthlyLimit: ut.PtrFloat(3000.50)}, nil},
		{"Missing not null field", &model.Card{CardID: uuid.New().String(), DailyLimit: ut.PtrFloat(100.50), MonthlyLimit: ut.PtrFloat(3000.50)}, &mysql.MySQLError{Number: 1048, Message: "Column 'name' cannot be null"}},
	}
	ctx := context.TODO()
	db := setupCardRepoDB(t)
	r := NewCard(db)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.Create(ctx, tt.card)

			assert.EqualValues(t, err, tt.error)
		})
	}
}

func TestGetCards(t *testing.T) {
	tests := []struct {
		name  string
		card  []*model.Card
		error error
	}{
		{"Valid card request", make([]*model.Card, 2), nil},
	}
	ctx := context.TODO()
	db := setupCardRepoDB(t)
	r := NewCard(db)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := r.Get(ctx)

			assert.EqualValues(t, err, tt.error)
		})
	}
}

func TestUpdateCard(t *testing.T) {
	tests := []struct {
		name    string
		cardId  string
		cardRes *model.Card
		error   error
	}{
		{"Card Update", "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada", &model.Card{Name: ut.PtrString("Testing name"), DailyLimit: ut.PtrFloat(100.50), MonthlyLimit: ut.PtrFloat(3000.50)}, nil},
	}
	ctx := context.TODO()
	db := setupCardRepoDB(t)
	r := NewCard(db)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.Update(ctx, tt.cardId, tt.cardRes)

			assert.EqualValues(t, err, tt.error)
		})
	}
}

func TestDeleteCard(t *testing.T) {
	tests := []struct {
		name   string
		cardId string
		error  error
	}{
		{"Card Delete", "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada", nil},
	}
	ctx := context.TODO()
	db := setupCardRepoDB(t)
	r := NewCard(db)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.Delete(ctx, tt.cardId)

			assert.EqualValues(t, err, tt.error)
		})
	}
}
