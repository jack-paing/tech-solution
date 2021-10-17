package card

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"tech-solution/model"
)

type Handler interface {
	Create(c *gin.Context)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(c *gin.Context) {
	var card *model.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		log.Errorf("JSON Bind Error %v", err)
		c.JSON(http.StatusBadRequest, "Invalid data model")
		return
	}

	err := h.service.Create(c.Request.Context(), card)
	if err != nil {
		log.Errorf("error in creating card %v", err)
		c.JSON(http.StatusBadRequest, "Error in creating card")
		return
	}
	c.JSON(http.StatusOK, "card is created")
}
