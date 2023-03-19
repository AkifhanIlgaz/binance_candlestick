package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/cinar/indicator"
	"github.com/joho/godotenv"
	"github.com/rodaine/table"
)

var APIKEY = os.Getenv("APIKEY")
var SECRETKEY = os.Getenv("SECRETKEY")
var CLIENT = binance.NewClient(APIKEY, SECRETKEY)

func main() {
	godotenv.Load(".env")

	// TODO: get symbol interval limit as an cli argument
	candleSticks := getCandleSticks("BTCUSDT", "1h", 1000)

	var (
		highs    []float64
		lows     []float64
		closings []float64
		volumes  []int64
	)

	for _, candle := range candleSticks[950:] {
		highs = append(highs, parseFloat(candle.High))
		lows = append(lows, parseFloat(candle.Low))
		closings = append(closings, parseFloat(candle.Close))
		volumes = append(volumes, parseInt(candle.Volume))
	}

	ema := indicator.Ema(20, closings)
	cmf := indicator.ChaikinMoneyFlow(highs, lows, closings, volumes)

	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	}

	tbl := table.New("Open Time", "Open Price", "Low", "High", "Closing Price", "EMA", "CMF")

	for i, candle := range candleSticks[950:] {

		tbl.AddRow(time.UnixMilli(candle.OpenTime).UTC(), candle.Open, candle.Low, candle.High, candle.Close, ema[i], cmf[i])
	}

	tbl.Print()
}

func getCandleSticks(symbol, interval string, limit int) []*binance.Kline {
	klineClient := CLIENT.NewKlinesService().Symbol(symbol).Interval(interval).Limit(limit)

	candleSticks, err := klineClient.Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return candleSticks
}

func parseFloat(s string) float64 {
	result, err := strconv.ParseFloat(s, 64)
	checkError(err)
	return result
}

func parseInt(s string) int64 {
	result, err := strconv.ParseFloat(s, 64)
	checkError(err)
	return int64(result)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
