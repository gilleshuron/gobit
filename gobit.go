package main

import (
	"github.com/gilleshuron/gobit/market"
	// "github.com/gilleshuron/gobit/model"
	"log"
	"os"
	"time"
)

func main() {
	LOG := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var (
		market market.Bitstamp
	)
	go market.Start()
	for {
		select {
		case trade := <-market.Trades:
			LOG.Printf("%s %s. Price:%.2f$ Amount:%.2fBTC Id:%d", trade.Market, trade.Move, trade.Price, trade.Amount, trade.Id)
		case <-market.Books:
			// LOG.Printf("%s Bids: %v Asks: %v ", book.Market, book.Bids, book.Asks)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
