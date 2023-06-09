package router

import (
	h "github.com/avcwisesa/card-dealer/handler"

	"github.com/gin-gonic/gin"
)

func New(h h.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", h.PingHandler())
	router.POST("/deck", h.NewDeckHandler())
	router.GET("/deck/:id", h.OpenDeckHandler())
	router.POST("/deck/:id/draw", h.DrawFromDeckHandler())

	return router
}
