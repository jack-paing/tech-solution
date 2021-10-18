package card

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"tech-solution/model"
	ut "tech-solution/utils"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		cardReq  *model.Card
		error    error
		wantCode int
	}{
		{
			"Valid Request", `
			{
				"name":"card name 2",
				"description": "Test card",
				"daily_limit": 100.50,
				"monthly_limit":3000
			}
			`, &model.Card{
				Name:         ut.PtrString("card name 2"),
				Desc:         ut.PtrString("Test card"),
				DailyLimit:   ut.PtrFloat(100.50),
				MonthlyLimit: ut.PtrFloat(3000),
			}, nil, http.StatusOK,
		},
		{
			"Valid Request Body", `
			{
				"name"
			}
			`, nil, nil, http.StatusBadRequest,
		},
	}
	serviceMock := &ServiceMock{}
	r := gin.Default()
	r.POST("/cards", NewHandler(serviceMock).Create)

	for _, tt := range tests {
		serviceMock.On("Create", context.TODO(), tt.cardReq).Return(tt.error)
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/cards", strings.NewReader(tt.body))
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, w.Code)
			}
		})
	}
}

func TestUpdateHandler(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		cardId   string
		cardReq  *model.Card
		error    error
		wantCode int
	}{
		{
			"Valid Request", `
			{
				"name":"card name update",
				"description": "Test card"
			}
			`, "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada",
			&model.Card{
				Name: ut.PtrString("card name update"),
				Desc: ut.PtrString("Test card"),
			}, nil, http.StatusOK,
		},
		{
			"Valid Request Body", `
			{
				"name"
			}
			`, "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada", nil, nil, http.StatusBadRequest,
		},
	}
	serviceMock := &ServiceMock{}
	r := gin.Default()
	r.PATCH("/cards/:cardId", NewHandler(serviceMock).Update)

	for _, tt := range tests {
		serviceMock.On("Update", context.TODO(), tt.cardId, tt.cardReq).Return(tt.error)
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PATCH", "/cards/"+tt.cardId, strings.NewReader(tt.body))
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, w.Code)
			}
		})
	}
}

func TestGetHandler(t *testing.T) {
	tests := []struct {
		name     string
		error    error
		wantCode int
	}{
		{"Get cards", nil, http.StatusOK},
	}
	serviceMock := &ServiceMock{}
	r := gin.Default()
	r.GET("/cards", NewHandler(serviceMock).Get)

	for _, tt := range tests {
		cards := make([]*model.Card, 1)
		serviceMock.On("Get", context.TODO()).Return(cards, tt.error)
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/cards", nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, w.Code)
			}
		})
	}
}

func TestDeleteHandler(t *testing.T) {
	tests := []struct {
		name     string
		cardId   string
		error    error
		wantCode int
	}{
		{"Valid id", "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada", nil, http.StatusOK},
		{"Invalid id", "585cbfc0-41dd-4e9f-b7aa-3f32a2b", errors.New("cannot find id"), http.StatusBadRequest},
	}
	serviceMock := &ServiceMock{}
	r := gin.Default()
	r.DELETE("/cards/:cardId", NewHandler(serviceMock).Delete)

	for _, tt := range tests {
		serviceMock.On("Delete", context.TODO(), tt.cardId).Return(tt.error)
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/cards/"+tt.cardId, nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, w.Code)
			}
		})
	}
}

func TestGetByCardIdHandler(t *testing.T) {
	tests := []struct {
		name     string
		param    string
		error    error
		wantCode int
	}{
		{"Valid id", "585cbfc0-41dd-4e9f-b7aa-3f32a2ba4ada", nil, http.StatusOK},
		{"Invalid id", "585cbfc0-41dd-4e9f-b7aa-3f32a2b", errors.New("cannot find id"), http.StatusBadRequest},
	}
	serviceMock := &ServiceMock{}
	r := gin.Default()
	r.GET("/cards/:cardId", NewHandler(serviceMock).GetByCardId)

	for _, tt := range tests {
		serviceMock.On("GetByCardId", context.TODO(), tt.param).Return(nil, tt.error)
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/cards/"+tt.param, nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, w.Code)
			}
		})
	}
}
