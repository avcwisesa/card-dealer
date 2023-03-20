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

func (suite *DealerTestSuite) TestCreateShuffledDeck() {
	dealer := domain.NewDealer()
	deck := dealer.CreateDeck(true)

	assert.NotNil(suite.T(), deck.GetID())
	assert.Equal(suite.T(), true, deck.IsShuffled())
	assert.Equal(suite.T(), 52, deck.CardsRemainingCount())
}

func (suite *DealerTestSuite) TestCreateCustomDeck() {
	dealer := domain.NewDealer()
	deck := dealer.CreateCustomDeck(true, "KS,KD,KC,KH")

	assert.NotNil(suite.T(), deck.GetID())
	assert.Equal(suite.T(), true, deck.IsShuffled())
	assert.Equal(suite.T(), 4, deck.CardsRemainingCount())
}

func (suite *DealerTestSuite) TestDrawFromDeck() {
	dealer := domain.NewDealer()

	deck := dealer.CreateDeck(false)
	card := dealer.DrawFromDeck(deck.GetID())

	assert.Equal(suite.T(), "AS", card.Content)
}

func TestDealerTestSuite(t *testing.T) {
	suite.Run(t, new(DealerTestSuite))
}
