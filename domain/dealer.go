package domain

import "strings"

type Dealer interface {
	CreateDeck(isShuffled bool) Deck
	CreateCustomDeck(isShuffled bool, cardsString string) Deck
}

type dealer struct{}

func (d *dealer) CreateDeck(isShuffled bool) Deck {
	return NewDeck(isShuffled, 52)
}

func (d *dealer) CreateCustomDeck(isShuffled bool, cardsString string) Deck {
	cards := strings.Split(cardsString, ",")

	return NewDeck(isShuffled, len(cards))
}

func NewDealer() Dealer {
	return &dealer{}
}
