package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"tech-solution/internal/card"
	"tech-solution/internal/config"
	"tech-solution/internal/health"
)

func main() {

	// load config
	config, err := config.Load()
	if err != nil {
		log.Errorf("Error in loading config %v", err)
	}
	db, err := getDb(config.DSN)
	if err != nil {
		fmt.Println(err)
		log.Error("Error in loading %v", err)
	}
	r := setupRouter(db)
	r.Run(fmt.Sprintf(":%v", config.Port))
}

func getDb(dsn string) (*sqlx.DB, error) {
	fmt.Println(dsn)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
func setupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", health.Check)

	cardService := card.NewService(db)
	cardHandler := card.NewHandler(cardService)
	r.POST("/cards", cardHandler.Create)
	r.GET("/cards", cardHandler.Get)
	r.GET("/cards/:cardId", cardHandler.GetByCardId)
	r.DELETE("/cards/:cardId", cardHandler.Delete)
	r.PATCH("/cards/:cardId", cardHandler.Update)

	return r
}
