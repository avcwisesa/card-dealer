package router_test

import (
	"github.com/avcwisesa/card-dealer/mocks"
	"github.com/avcwisesa/card-dealer/router"

	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RouterTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *RouterTestSuite) SetupTest() {
	mockHandler := mocks.NewMockHandler()
	suite.router = router.New(mockHandler)
}

func (suite *RouterTestSuite) TestPingRoute() {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "pong", w.Body.String())
}

func (suite *RouterTestSuite) TestNewDeckRoute() {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/deck", nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), "ok", w.Body.String())
}

func (suite *RouterTestSuite) TestOpenDeckRoute() {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/deck/5bde8679-2884-4eee-b572-38673b11c9bf", nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "ok", w.Body.String())
}

func (suite *RouterTestSuite) TestDrawFromDeckRoute() {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/deck/5bde8679-2884-4eee-b572-38673b11c9bf/draw", nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "ok", w.Body.String())
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}
