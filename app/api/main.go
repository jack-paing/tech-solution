package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"tech-solution/internal/config"
	"tech-solution/internal/health"
)

func main() {
	r := setupRouter()
	// load config
	config, err := config.Load()
	if err != nil {
		log.Error("Error in loading config")
	}
	r.Run(fmt.Sprintf(":%v", config.Port))
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", health.Check)

	return r
}
