package market

import (
	"encoding/json"
	"github.com/ehq/pusher-go"
	"github.com/gilleshuron/gobit/model"
	"log"
	"os"
	"time"
)

const (
	name string = "Bitstamp"
)

type Bitstamp struct {
	trades         chan model.Trade
	books          chan model.Book
	tradingHistory []model.Trade
}

func NewBitstamp() *Bitstamp {
	return &Bitstamp{
		trades: make(chan model.Trade),
		books:  make(chan model.Book),
	}
}

func (b *Bitstamp) TradingHistory() []model.Trade {
	return b.tradingHistory
}

func (b *Bitstamp) Trades() chan model.Trade {
	return b.trades
}

func (b *Bitstamp) Books() chan model.Book {
	return b.books
}

func (b *Bitstamp) Name() string {
	return name
}

func (b *Bitstamp) Run() {
	LOG := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	b.tradingHistory = TradingHistory()

	client, err := pusher.Connect("de504dc5763aeef9ff52")

	if err != nil {
		LOG.Fatal(err)
		return
	}
	LOG.Printf("%s Connected", name)

	// previousTrade := model.Trade{}

	client.Subscribe("live_trades")
	client.Subscribe("order_book")

	client.On("trade", func(data string) {
		// LOG.Println(data)
		trade := model.Trade{}
		err := json.Unmarshal([]byte(data), &trade)
		if err != nil {
			LOG.Fatal(err)
		}
		trade.Market = name
		trade.Date = time.Now().UnixNano() / 1000000
		// if previousTrade.Id == 0 {
		// 	trade.Move = "?"
		// } else if trade.Price > previousTrade.Price {
		// 	trade.Move = "UP"
		// } else if trade.Price == previousTrade.Price {
		// 	trade.Move = "-"
		// } else {
		// 	trade.Move = "DOWN"
		// }
		b.trades <- trade
		b.tradingHistory = append(b.tradingHistory, trade)
		// previousTrade = trade
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
		b.books <- book
	})

}
