package router

import (
	mocks "github.com/avcwisesa/card-dealer/mocks"

	"net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
	"github.com/gin-gonic/gin"
)

type RouterTestSuite struct {
    suite.Suite
    router *gin.Engine
}

func (suite *RouterTestSuite) SetupTest() {
	mockHandler := mocks.NewMockHandler()
	suite.router = New(mockHandler)
}


func (suite *RouterTestSuite) TestPingRoute() {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/ping", nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "pong", w.Body.String())
}

func TestRouterTestSuite(t *testing.T) {
    suite.Run(t, new(RouterTestSuite))
}
