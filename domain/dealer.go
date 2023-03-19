package domain

type Dealer interface {
	CreateDeck() Deck
}

type dealer struct{}

func (d *dealer) CreateDeck() Deck {
	return NewDeck(false, 52)
}

func NewDealer() Dealer {
	return &dealer{}
}
