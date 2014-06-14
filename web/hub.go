package web

import (
	"encoding/json"
	"github.com/gilleshuron/gobit/model"
	"log"
	"os"
	"time"
)

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Registered markets
	markets map[model.Market]bool

	// Inbound messages from the connections.
	Broadcast chan []byte

	// Register Markets
	RegisterMarket chan model.Market

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

var h = Hub{
	Broadcast:      make(chan []byte),
	RegisterMarket: make(chan model.Market),
	register:       make(chan *connection),
	unregister:     make(chan *connection),
	connections:    make(map[*connection]bool),
	markets:        make(map[model.Market]bool),
}

var LOG *log.Logger

func init() {
	LOG = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}

func H() Hub {
	return h
}

func (h *Hub) marketPump() {
	for {
		for m := range h.markets {
			select {
			case trade := <-m.Trades():
				LOG.Printf("%s. Price:%.2f$ Amount:%.2fBTC Id:%d", trade.Market, trade.Price, trade.Amount, trade.Id)
				j, _ := json.Marshal(trade)
				h.Broadcast <- j
			case <-m.Books():
			default:
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (h *Hub) Run() {
	go h.marketPump()
	for {
		select {
		case m := <-h.RegisterMarket:
			LOG.Printf("%s Registered", m.Name())
			h.markets[m] = true
		case msg := <-h.Broadcast:
			for c := range h.connections {
				select {
				case c.send <- msg:
				default:
					close(c.send)
					delete(h.connections, c)
				}
			}
		case c := <-h.register:
			h.connections[c] = true
			for m := range h.markets {
				// History
				for _, t := range m.TradingHistory() {
					j, _ := json.Marshal(t)
					select {
					case c.send <- j:
					default:
						time.Sleep(10 * time.Millisecond)
					}
				}
			}
		case c := <-h.unregister:
			delete(h.connections, c)
			close(c.send)
		}
	}
}
