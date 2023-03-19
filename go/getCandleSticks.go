package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
)

var APIKEY = os.Getenv("APIKEY")
var SECRETKEY = os.Getenv("SECRETKEY")
var CLIENT = binance.NewClient(APIKEY, SECRETKEY)

func getCandleSticks(symbol, interval string, limit int) []*binance.Kline {
	klineClient := CLIENT.NewKlinesService().Symbol(symbol).Interval(interval).Limit(limit)

	candleSticks, err := klineClient.Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return candleSticks
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseArgs() (*string, *string, *int) {
	symbol := flag.String("pair", "", "name of the pair")
	interval := flag.String("interval", "4h", "interval")
	limit := flag.Int("limit", 500, "number of candles")

	flag.Parse()

	return symbol, interval, limit
}
