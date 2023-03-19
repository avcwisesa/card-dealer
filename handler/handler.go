package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	PingHandler() func(c *gin.Context)
}

type handler struct{}

func (h *handler) PingHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func New() Handler {
	return &handler{}
}
