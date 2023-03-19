package handler

import (
	"github.com/avcwisesa/card-dealer/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	PingHandler() func(c *gin.Context)
	NewDeckHandler() func(c *gin.Context)
}

type handler struct {
	dealer domain.Dealer
}

func (h *handler) PingHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func (h *handler) NewDeckHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		newDeck := h.dealer.CreateDeck()

		c.JSON(http.StatusCreated, gin.H{
			"deck_id":   newDeck.GetID(),
			"shuffled":  newDeck.IsShuffled(),
			"remaining": newDeck.CardsRemaining(),
		})
	}
}

func New(dealer domain.Dealer) Handler {
	return &handler{dealer: dealer}
}
