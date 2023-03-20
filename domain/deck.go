package domain

import "github.com/google/uuid"

type Deck interface {
	GetID() string
	IsShuffled() bool
	CardsRemainingCount() int
	CardsRemaining() []Card
}

type deck struct {
	ID         string
	isShuffled bool
	cards      []Card
}

func (d *deck) GetID() string {
	return d.ID
}

func (d *deck) IsShuffled() bool {
	return d.isShuffled
}

func (d *deck) CardsRemainingCount() int {
	return len(d.cards)
}

func (d *deck) CardsRemaining() []Card {
	return d.cards
}

func NewDeck(isShuffled bool, cards []Card) *deck {
	return &deck{
		ID:         uuid.NewString(),
		isShuffled: isShuffled,
		cards:      cards,
	}
}
