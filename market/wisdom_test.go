package market

import (
	"log"
	"testing"
)

func TestTradingHistory(t *testing.T) {
	trds := TradingHistory()
	log.Println(trds)
}
