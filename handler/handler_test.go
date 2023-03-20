package handler_test

import (
	"github.com/avcwisesa/card-dealer/domain"
	"github.com/avcwisesa/card-dealer/handler"
	"github.com/avcwisesa/card-dealer/mocks"

	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	handler handler.Handler
	dealer  *mocks.MockDealer
}

func (suite *HandlerTestSuite) SetupTest() {
	dealer := mocks.NewDealer()
	suite.handler = handler.New(dealer)
}

func (suite *HandlerTestSuite) TestPing() {
	pingHandler := suite.handler.PingHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	pingHandler(c)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "{\"status\":\"ok\"}", w.Body.String())
}

func (suite *HandlerTestSuite) TestDefaultNewDeck() {
	expectedDeckID := "5bde8679-2884-4eee-b572-38673b11c9bf"
	expectedRemaining := 52
	expectedIsShuffled := false
	expectedDeck := mocks.NewDeck(
		expectedDeckID,
		expectedIsShuffled,
		expectedRemaining,
	)
	expectedDeck.On("CardsRemainingCount").Return(52)

	dealer := mocks.NewDealer()
	dealer.On("CreateDeck", false).Return(expectedDeck)

	handler := handler.New(dealer)
	newDeckHandler := handler.NewDeckHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	newDeckHandler(c)

	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), "{\"deck_id\":\"5bde8679-2884-4eee-b572-38673b11c9bf\",\"remaining\":52,\"shuffled\":false}", w.Body.String())
}

func (suite *HandlerTestSuite) TestShuffledNewDeck() {
	expectedDeckID := "5bde8679-2884-4eee-b572-38673b11c9bf"
	expectedRemaining := 52
	expectedIsShuffled := true
	expectedDeck := mocks.NewDeck(
		expectedDeckID,
		expectedIsShuffled,
		expectedRemaining,
	)
	expectedDeck.On("CardsRemainingCount").Return(52)

	dealer := mocks.NewDealer()
	dealer.On("CreateDeck", true).Return(expectedDeck)

	handler := handler.New(dealer)
	newDeckHandler := handler.NewDeckHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest("POST", "/deck?is_shuffled=true", nil)
	c.Request = req

	newDeckHandler(c)

	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), "{\"deck_id\":\"5bde8679-2884-4eee-b572-38673b11c9bf\",\"remaining\":52,\"shuffled\":true}", w.Body.String())
}

func (suite *HandlerTestSuite) TestNewCustomDeck() {
	expectedDeckID := "5bde8679-2884-4eee-b572-38673b11c9bf"
	expectedRemaining := 4
	expectedIsShuffled := true
	expectedDeck := mocks.NewDeck(
		expectedDeckID,
		expectedIsShuffled,
		expectedRemaining,
	)
	expectedDeck.On("CardsRemainingCount").Return(4)

	dealer := mocks.NewDealer()
	dealer.On("CreateCustomDeck", true, "KS,KD,KC,KH").Return(expectedDeck)

	handler := handler.New(dealer)
	newDeckHandler := handler.NewDeckHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest("POST", "/deck?is_shuffled=true&cards=KS,KD,KC,KH", nil)
	c.Request = req

	newDeckHandler(c)

	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), "{\"deck_id\":\"5bde8679-2884-4eee-b572-38673b11c9bf\",\"remaining\":4,\"shuffled\":true}", w.Body.String())
}

func (suite *HandlerTestSuite) TestOpenUnavailableDeck() {
	dealer := mocks.NewDealer()
	dealer.On("GetDeck", "5bde8679-2884-4eee-b572-38673b11c9bf").Return(nil)

	handler := handler.New(dealer)
	openDeckHandler := handler.OpenDeckHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{gin.Param{Key: "id", Value: "5bde8679-2884-4eee-b572-38673b11c9bf"}}

	openDeckHandler(c)

	assert.Equal(suite.T(), 404, w.Code)
	assert.Equal(suite.T(), "{\"error\":\"Deck not available\"}", w.Body.String())
}

func (suite *HandlerTestSuite) TestOpenDeck() {
	expectedDeckID := "5bde8679-2884-4eee-b572-38673b11c9bf"
	expectedDeck := mocks.NewCustomDeck(
		expectedDeckID,
		false,
		[]domain.Card{
			domain.Card{Content: "AH"},
			domain.Card{Content: "7C"},
		},
	)
	expectedDeck.On("CardsRemainingCount").Return(2)

	dealer := mocks.NewDealer()
	dealer.On("GetDeck", expectedDeckID).Return(expectedDeck)

	handler := handler.New(dealer)
	openDeckHandler := handler.OpenDeckHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{gin.Param{Key: "id", Value: expectedDeckID}}

	openDeckHandler(c)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "{\"cards\":[{\"code\":\"AH\",\"suite\":\"HEARTS\",\"value\":\"ACE\"},{\"code\":\"7C\",\"suite\":\"CLUBS\",\"value\":\"7\"}],\"deck_id\":\"5bde8679-2884-4eee-b572-38673b11c9bf\",\"remaining\":2,\"shuffled\":false}", w.Body.String())
}

func (suite *HandlerTestSuite) TestDrawFromUnavailableDeck() {
	deckID := "5bde8679-2884-4eee-b572-38673b11c9bf"
	expectedCard := domain.Card{Content: "test"}

	dealer := mocks.NewDealer()
	dealer.On("DrawFromDeck", deckID, 1).Return(expectedCard)

	handler := handler.New(dealer)
	drawFromDeck := handler.DrawFromDeckHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{gin.Param{Key: "id", Value: deckID}}

	drawFromDeck(c)

	assert.Equal(suite.T(), 404, w.Code)
	assert.Equal(suite.T(), "{\"error\":\"Deck not available\"}", w.Body.String())
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
