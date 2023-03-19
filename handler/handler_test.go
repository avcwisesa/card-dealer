package handler

import (
	"net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
	"github.com/gin-gonic/gin"
)

type HandlerTestSuite struct {
    suite.Suite
	handler Handler
}

func (suite *HandlerTestSuite) SetupTest() {
	suite.handler = New()
}


func (suite *HandlerTestSuite) TestPingHandler() {
	pingHandler := suite.handler.PingHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	pingHandler(c)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "{\"status\":\"ok\"}", w.Body.String())
}

func TestHandlerTestSuite(t *testing.T) {
    suite.Run(t, new(HandlerTestSuite))
}
