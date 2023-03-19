package main

import (
	"github.com/avcwisesa/card-dealer/domain"
	"github.com/avcwisesa/card-dealer/handler"
	"github.com/avcwisesa/card-dealer/router"
)

func main() {
	dealer := domain.NewDealer()
	h := handler.New(dealer)
	r := router.New(h)
	r.Run()
}
