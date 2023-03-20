package handler

import (
	"github.com/avcwisesa/card-dealer/domain"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	PingHandler() func(c *gin.Context)
	NewDeckHandler() func(c *gin.Context)
	OpenDeckHandler() func(c *gin.Context)
	DrawFromDeckHandler() func(c *gin.Context)
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
		isShuffled, err := strconv.ParseBool(c.Query("is_shuffled"))
		if err != nil {
			isShuffled = false
		}

		var newDeck domain.Deck
		cards := c.Query("cards")
		if cards == "" {
			newDeck = h.dealer.CreateDeck(isShuffled)
		} else {
			newDeck = h.dealer.CreateCustomDeck(isShuffled, cards)
		}

		c.JSON(http.StatusCreated, gin.H{
			"deck_id":   newDeck.GetID(),
			"shuffled":  newDeck.IsShuffled(),
			"remaining": newDeck.CardsRemainingCount(),
		})
	}
}

func (h *handler) OpenDeckHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		deckID := c.Param("id")

		deck := h.dealer.GetDeck(deckID)
		if deck == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Deck not available",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"deck_id":   deck.GetID(),
			"shuffled":  deck.IsShuffled(),
			"remaining": deck.CardsRemainingCount(),
			"cards":     domain.PresentPokerCards(deck.CardsRemaining()),
		})
	}
}

func (h *handler) DrawFromDeckHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		deckID := c.Param("id")

		card := h.dealer.DrawFromDeck(deckID)
		if card.IsEmpty() {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Deck not available or ran out of cards",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"cards": domain.PresentPokerCards([]domain.Card{card}),
		})
	}
}

func New(dealer domain.Dealer) Handler {
	return &handler{dealer: dealer}
}
