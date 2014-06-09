package main

import (
    "github.com/ehq/pusher-go"
    "log"
    "os"
    "encoding/json"
)   
    

type Trade struct {
     Price float64 `json:"price"`
     Amount float64 `json:"amount"`
     Id uint64 `json:"id"`
} 


func main() {
    LOG := log.New(os.Stdout,"BITSTAMP ",log.Ldate|log.Ltime) 

    client, err := pusher.Connect("de504dc5763aeef9ff52")

    if err != nil {
        LOG.Fatal(err)
        return
    }
    LOG.Printf("Connected")

    client.Subscribe("live_trades")

    trades := make(chan Trade)

    client.On("trade", func(data string) {
        // LOG.Println(data)
        trade := Trade{}
        err := json.Unmarshal([]byte(data), &trade)
        if (err != nil) {
           LOG.Println(err)
           return
        }    
        trades <- trade
    })
    var previousTrade Trade
    for {
        trade := <-trades
        var trend string
        if previousTrade.Id==0 {
            trend="?"
        } else if trade.Price>previousTrade.Price {
            trend="UP"
        } else if trade.Price==previousTrade.Price {
            trend="-"
        } else {
            trend="DOWN"
        }    
        LOG.Printf("%s. Price:%.2f$ Amount:%.2fBTC Id:%d", trend, trade.Price, trade.Amount, trade.Id)
        previousTrade=trade
    }

}