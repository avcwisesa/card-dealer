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

func (h *handlerMock) NewDeckHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusCreated, "ok")
	}
}

func (h *handlerMock) OpenDeckHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}
}

func (h *handlerMock) DrawFromDeckHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	}
}

func NewMockHandler() handler.Handler {
	return &handlerMock{}
}
