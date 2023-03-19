package domain

type Dealer interface {
	CreateDeck(isShuffled bool) Deck
}

type dealer struct{}

func (d *dealer) CreateDeck(isShuffled bool) Deck {
	return NewDeck(isShuffled, 52)
}

func NewDealer() Dealer {
	return &dealer{}
}
