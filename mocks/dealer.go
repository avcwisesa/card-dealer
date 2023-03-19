package mocks

import (
	"github.com/avcwisesa/card-dealer/domain"

	"github.com/stretchr/testify/mock"
)

type MockDealer struct {
	mock.Mock
}

func (d *MockDealer) CreateDeck(isShuffled bool) domain.Deck {
	args := d.Called(isShuffled)
	return args.Get(0).(domain.Deck)
}

func NewDealer() *MockDealer {
	return &MockDealer{}
}
