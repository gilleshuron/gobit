package market

import (
	"encoding/json"
	"github.com/ehq/pusher-go"
	"github.com/gilleshuron/gobit/model"
	"log"
	"os"
)

const (
	name string = "Bitstamp"
)

type Bitstamp struct {
	Trades chan model.Trade
	Books  chan model.Book
}

func NewBitstamp() Bitstamp {
	return Bitstamp{
		Trades: make(chan model.Trade),
		Books:  make(chan model.Book),
	}
}

func (b *Bitstamp) Run() {
	LOG := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	client, err := pusher.Connect("de504dc5763aeef9ff52")

	if err != nil {
		LOG.Fatal(err)
		return
	}
	LOG.Printf("%s Connected", name)

	previousTrade := model.Trade{}

	client.Subscribe("live_trades")
	client.Subscribe("order_book")

	client.On("trade", func(data string) {
		// LOG.Println(data)
		trade := model.Trade{}
		err := json.Unmarshal([]byte(data), &trade)
		if err != nil {
			LOG.Println(err)
			return
		}
		trade.Market = name
		if previousTrade.Id == 0 {
			trade.Move = "?"
		} else if trade.Price > previousTrade.Price {
			trade.Move = "UP"
		} else if trade.Price == previousTrade.Price {
			trade.Move = "-"
		} else {
			trade.Move = "DOWN"
		}
		b.Trades <- trade
		previousTrade = trade
	})

	client.On("data", func(data string) {
		// LOG.Println(data)
		book := model.Book{}
		err := json.Unmarshal([]byte(data), &book)
		if err != nil {
			LOG.Println(err)
			return
		}
		book.Market = name
		b.Books <- book
	})

}
