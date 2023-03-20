package domain

import (
	"strconv"
	"strings"
)

type Dealer interface {
	CreateDeck(isShuffled bool) Deck
	CreateCustomDeck(isShuffled bool, cardsString string) Deck
	GetDeck(id string) Deck
}

type dealer struct {
	store *Store
}

type cardPresentation map[string]string

func (d *dealer) CreateDeck(isShuffled bool) Deck {
	deck := NewDeck(isShuffled, generateStandardPokerDeck())
	d.store.AddOrUpdate(deck)

	return deck
}

func (d *dealer) CreateCustomDeck(isShuffled bool, cardsString string) Deck {
	cardContents := strings.Split(cardsString, ",")
	cards := []Card{}
	for _, content := range cardContents {
		cards = append(cards, Card{Content: content})
	}

	return NewDeck(isShuffled, cards)
}

func (d *dealer) GetDeck(id string) Deck {
	return d.store.Get(id)
}

func generateStandardPokerDeck() []Card {
	cards := []Card{}
	suites := []string{"S", "D", "C", "H"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suite := range suites {
		for _, value := range values {
			card := Card{Content: value + suite}
			cards = append(cards, card)
		}
	}

	return cards
}

func PresentPokerDeck(deck Deck) []cardPresentation {
	deckPresentation := []cardPresentation{}

	valueMapping := map[string]string{
		"A": "ACE",
		"J": "JACK",
		"Q": "QUEEN",
		"K": "KING",
	}
	for i := 2; i <= 10; i++ {
		strValue := strconv.Itoa(i)
		valueMapping[strValue] = strValue
	}

	suiteMapping := map[string]string{
		"S": "SPADES",
		"D": "DIAMONDS",
		"C": "CLUBS",
		"H": "HEARTS",
	}

	for _, card := range deck.CardsRemaining() {
		value := card.Content[0 : len(card.Content)-1]
		suite := card.Content[len(card.Content)-1 : len(card.Content)]

		cp := cardPresentation{
			"value": valueMapping[value],
			"suite": suiteMapping[suite],
			"code":  card.Content,
		}
		deckPresentation = append(deckPresentation, cp)
	}

	return deckPresentation
}

func NewDealer() Dealer {
	return &dealer{store: NewStore()}
}
