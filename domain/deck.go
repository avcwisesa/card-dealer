package domain

import "github.com/google/uuid"

type Deck interface {
	GetID() string
	IsShuffled() bool
	CardsRemaining() int
}

type deck struct {
	deckID     string
	isShuffled bool
	remaining  int
}

func (d *deck) GetID() string {
	return d.deckID
}

func (d *deck) IsShuffled() bool {
	return d.isShuffled
}

func (d *deck) CardsRemaining() int {
	return d.remaining
}

func NewDeck(isShuffled bool, remaining int) Deck {
	return &deck{
		deckID:     uuid.NewString(),
		isShuffled: isShuffled,
		remaining:  remaining,
	}
}
