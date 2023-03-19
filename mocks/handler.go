package mocks

import (
	"github.com/avcwisesa/card-dealer/handler"

	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerMock struct{}

func (h *handlerMock) PingHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}

func NewMockHandler() handler.Handler {
	return &handlerMock{}
}
