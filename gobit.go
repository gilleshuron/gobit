package main

import (
	"github.com/gilleshuron/gobit/market"
	"github.com/gilleshuron/gobit/web"

	"github.com/gilleshuron/gobit/model"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	w := web.NewWeb()

	h := web.H()
	go h.Run()

	var m model.Market
	m = market.NewBitstamp()
	h.RegisterMarket <- m
	go m.Run()

	go w.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		panic("Show me the stack:")
		os.Exit(1)
	}()

	for {
		select {
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
