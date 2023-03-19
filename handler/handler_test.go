package handler_test

import (
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

func (suite *HandlerTestSuite) TestPingHandler() {
	pingHandler := suite.handler.PingHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	pingHandler(c)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "{\"status\":\"ok\"}", w.Body.String())
}

func (suite *HandlerTestSuite) TestNewDeckHandler() {
	expectedDeckID := "5bde8679-2884-4eee-b572-38673b11c9bf"
	expectedRemaining := 52
	expectedIsShuffled := false
	expectedDeck := mocks.NewDeck(
		expectedDeckID,
		expectedIsShuffled,
		expectedRemaining,
	)

	dealer := mocks.NewDealer()
	dealer.On("CreateDeck").Return(expectedDeck)

	handler := handler.New(dealer)
	newDeckHandler := handler.NewDeckHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	newDeckHandler(c)

	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), "{\"deck_id\":\"5bde8679-2884-4eee-b572-38673b11c9bf\",\"remaining\":52,\"shuffled\":false}", w.Body.String())
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
