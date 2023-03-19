package mocks

import (
	"github.com/avcwisesa/card-dealer/domain"

	"github.com/stretchr/testify/mock"
)

type MockDealer struct {
	mock.Mock
}

func (d *MockDealer) CreateDeck() domain.Deck {
	args := d.Called()
	return args.Get(0).(domain.Deck)
}

func NewDealer() *MockDealer {
	return &MockDealer{}
}
