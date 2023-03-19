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
	assert.Equal(suite.T(), 52, deck.CardsRemaining())
}

func (suite *DealerTestSuite) TestCreateShuffledDeck() {
	dealer := domain.NewDealer()
	deck := dealer.CreateDeck(true)

	assert.NotNil(suite.T(), deck.GetID())
	assert.Equal(suite.T(), true, deck.IsShuffled())
	assert.Equal(suite.T(), 52, deck.CardsRemaining())
}

func TestDealerTestSuite(t *testing.T) {
	suite.Run(t, new(DealerTestSuite))
}
