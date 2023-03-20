package domain_test

import (
	"github.com/avcwisesa/card-dealer/domain"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DealerTestSuite struct {
	suite.Suite
}

func (suite *DealerTestSuite) TestCreateDefaultDeck() {
	dealer := domain.NewDealer()
	deck := dealer.CreateDeck(false)

	assert.NotNil(suite.T(), deck.GetID())
	assert.Equal(suite.T(), false, deck.IsShuffled())
	assert.Equal(suite.T(), 52, deck.CardsRemainingCount())
}

func (suite *DealerTestSuite) TestCreateCustomDeck() {
	dealer := domain.NewDealer()
	deck := dealer.CreateCustomDeck(true, "KS,KD,KC,KH")

	assert.NotNil(suite.T(), deck.GetID())
	assert.Equal(suite.T(), true, deck.IsShuffled())
	assert.Equal(suite.T(), 4, deck.CardsRemainingCount())
}

func (suite *DealerTestSuite) TestCreateShuffledDeck() {
	dealer := domain.NewDealer()
	deck := dealer.CreateCustomDeck(true, "KS,KD,KC,KH")

	baseCardOrder := []string{"KS", "KD", "KC", "KH"}

	cards := deck.CardsRemaining()
	currentCardOrder := []string{}
	for _, card := range cards {
		currentCardOrder = append(currentCardOrder, card.Content)
	}

	assert.ElementsMatch(suite.T(), baseCardOrder, currentCardOrder)
	assert.NotEqualValues(suite.T(), baseCardOrder, currentCardOrder)
}

func (suite *DealerTestSuite) TestGetCreatedCustomDeck() {
	dealer := domain.NewDealer()
	deck := dealer.CreateCustomDeck(true, "KS,KD,KC,KH")

	retrievedDeck := dealer.GetDeck(deck.GetID())

	assert.Equal(suite.T(), retrievedDeck.GetID(), deck.GetID())
	assert.Equal(suite.T(), retrievedDeck.IsShuffled(), deck.IsShuffled())
	assert.Equal(suite.T(), retrievedDeck.CardsRemainingCount(), deck.CardsRemainingCount())
}

func (suite *DealerTestSuite) TestDrawFromDeck() {
	dealer := domain.NewDealer()

	deck := dealer.CreateDeck(false)
	card := dealer.DrawFromDeck(deck.GetID(), 1)[0]

	assert.Equal(suite.T(), "AS", card.Content)
}

func (suite *DealerTestSuite) TestDrawMultipleFromDeck() {
	dealer := domain.NewDealer()

	deck := dealer.CreateDeck(false)
	cards := dealer.DrawFromDeck(deck.GetID(), 3)

	assert.Equal(suite.T(), "AS", cards[0].Content)
	assert.Equal(suite.T(), "2S", cards[1].Content)
	assert.Equal(suite.T(), "3S", cards[2].Content)
}

func (suite *DealerTestSuite) TestDrawFromEmptyDeck() {
	dealer := domain.NewDealer()

	deck := dealer.CreateCustomDeck(false, "")
	cards := dealer.DrawFromDeck(deck.GetID(), 1)

	assert.Equal(suite.T(), 0, len(cards))
}

func (suite *DealerTestSuite) TestDrawTwoFromOneCardDeck() {
	dealer := domain.NewDealer()

	deck := dealer.CreateCustomDeck(false, "2S")
	cards := dealer.DrawFromDeck(deck.GetID(), 2)

	assert.Equal(suite.T(), len(cards), 1)
	assert.Equal(suite.T(), "2S", cards[0].Content)
}

func (suite *DealerTestSuite) TestDrawTwiceFromDeck() {
	dealer := domain.NewDealer()

	deck := dealer.CreateDeck(false)

	cards := dealer.DrawFromDeck(deck.GetID(), 1)
	assert.Equal(suite.T(), "AS", cards[0].Content)

	cards = dealer.DrawFromDeck(deck.GetID(), 1)
	assert.Equal(suite.T(), "2S", cards[0].Content)
}

func TestDealerTestSuite(t *testing.T) {
	suite.Run(t, new(DealerTestSuite))
}
