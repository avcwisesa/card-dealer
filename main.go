package main

import (
    "github.com/avcwisesa/card-dealer/handler"
    "github.com/avcwisesa/card-dealer/router"
)

func main() {
    h := handler.New()
	r := router.New(h)
    r.Run()
}
