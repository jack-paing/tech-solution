package card

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"tech-solution/model"
)

type Handler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	GetByCardId(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, "Invalid request body")
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

func (h *handler) Get(c *gin.Context) {
	cards, err := h.service.Get(c.Request.Context())
	if err != nil {
		log.Errorf("error in getting cards %v", err)
		c.JSON(http.StatusBadRequest, "Error in getting cards")
		return
	}
	c.JSON(http.StatusOK, cards)
}

func (h *handler) GetByCardId(c *gin.Context) {
	cardId := c.Param("cardId")
	card, err := h.service.GetByCardId(c.Request.Context(), cardId)
	if err != nil {
		log.Errorf("error in getting card %v", err)
		c.JSON(http.StatusBadRequest, "Error in getting card")
		return
	}
	c.JSON(http.StatusOK, card)
}

func (h *handler) Delete(c *gin.Context) {
	cardId := c.Param("cardId")
	err := h.service.Delete(c.Request.Context(), cardId)
	if err != nil {
		log.Errorf("error in deleting card %v", err)
		c.JSON(http.StatusBadRequest, "Error in deleting card")
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *handler) Update(c *gin.Context) {
	cardId := c.Param("cardId")
	var card *model.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		log.Errorf("JSON Bind Error %v", err)
		c.JSON(http.StatusBadRequest, "Invalid data model")
		return
	}
	err := h.service.Update(c.Request.Context(), cardId, card)
	if err != nil {
		log.Errorf("error in updating card %v", err)
		c.JSON(http.StatusBadRequest, "Error in updating card")
		return
	}
	c.JSON(http.StatusOK, "OK")
}
