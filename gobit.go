package main

import (
	"encoding/json"
	"github.com/gilleshuron/gobit/market"
	"github.com/gilleshuron/gobit/web"

	// "github.com/gilleshuron/gobit/model"
	"log"
	"os"
	"time"
)

func main() {
	LOG := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	m := market.NewBitstamp()
	w := web.NewWeb()
	h := web.H()
	go h.Run()
	go m.Run()
	go w.Run()
	for {
		select {
		case trade := <-m.Trades:
			LOG.Printf("%s %s. Price:%.2f$ Amount:%.2fBTC Id:%d", trade.Market, trade.Move, trade.Price, trade.Amount, trade.Id)
			j, _ := json.Marshal(trade)
			h.Broadcast <- j
		case <-m.Books:
			// LOG.Printf("%s Bids: %v Asks: %v ", book.Market, book.Bids, book.Asks)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
