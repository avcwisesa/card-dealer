package mocks

import (
	"github.com/avcwisesa/card-dealer/domain"

	"github.com/stretchr/testify/mock"
)

type MockDeck struct {
	mock.Mock
	deckID     string
	isShuffled bool
	remaining  int
	cards      []domain.Card
}

func (d *MockDeck) GetID() string {
	return d.deckID
}

func (d *MockDeck) IsShuffled() bool {
	return d.isShuffled
}

func (d *MockDeck) CardsRemainingCount() int {
	args := d.Called()
	return args.Get(0).(int)
}

func (d *MockDeck) CardsRemaining() []domain.Card {
	return d.cards
}

func (d *MockDeck) SetCards(cards []domain.Card) {
	d.Called(cards)
}

func NewDeck(id string, isShuffled bool, remaining int) *MockDeck {
	return &MockDeck{
		deckID:     id,
		isShuffled: isShuffled,
		remaining:  remaining,
	}
}

func NewCustomDeck(id string, isShuffled bool, cards []domain.Card) *MockDeck {
	return &MockDeck{
		deckID:     id,
		isShuffled: isShuffled,
		remaining:  len(cards),
		cards:      cards,
	}
}
