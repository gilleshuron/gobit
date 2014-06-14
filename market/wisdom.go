// curl 'https://s2.bitcoinwisdom.com/trades?since=0&sid=e93ca5&symbol=bitstampbtcusd&nonce=1402583233492'
// -H 'Origin: https://bitcoinwisdom.com' -H 'Accept-Encoding: gzip,deflate,sdch' -H 'Accept-Language: en-US,en;q=0.8,fr;q=0.6'
// -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.114 Safari/537.36'
// -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://bitcoinwisdom.com/'
// -H 'Connection: keep-alive' -H 'Cache-Control: max-age=0' --compressed

package market

import (
	"encoding/json"
	// "fmt"
	"github.com/gilleshuron/gobit/model"
	"io/ioutil"
	"log"
	"net/http"
	// "time"
)

type WTrade struct {
	Tid    uint64
	Date   int64
	Price  float64
	Amount float64
}

func TradingHistory() (trades []model.Trade) {
	log.Println("Getting wisdom history...")
	trades = make([]model.Trade, 100)

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	// t := fmt.Sprint(time.Now().Unix())
	resp, err := client.Get("https://s2.bitcoinwisdom.com/trades?since=0&sid=e93ca5&symbol=bitstampbtcusd")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var wtrades []WTrade
	json.Unmarshal(body, &wtrades)

	trades = make([]model.Trade, len(wtrades))

	for i, wt := range wtrades {
		if wt.Tid != 0 {
			trade := model.Trade{Id: wt.Tid, Price: wt.Price, Amount: wt.Amount, Market: "Bitstamp", Date: wt.Date * 1000}
			trades[i] = trade
		}
	}
	return
}
