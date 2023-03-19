package mocks

import "github.com/avcwisesa/card-dealer/domain"

type mockDeck struct {
	deckID     string
	isShuffled bool
	remaining  int
}

func (d *mockDeck) GetID() string {
	return d.deckID
}

func (d *mockDeck) IsShuffled() bool {
	return d.isShuffled
}

func (d *mockDeck) CardsRemaining() int {
	return d.remaining
}

func NewDeck(id string, isShuffled bool, remaining int) domain.Deck {
	return &mockDeck{
		deckID:     id,
		isShuffled: isShuffled,
		remaining:  remaining,
	}
}
