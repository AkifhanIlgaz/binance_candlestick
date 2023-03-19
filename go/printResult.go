package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/rodaine/table"
)

func printResult(candleSticks []*binance.Kline, ema, cmf *[]float64) {

	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return strings.ToUpper(fmt.Sprintf(format, vals...))
	}

	tbl := table.New("Open Time", "Open Price", "Low", "High", "Closing Price", "EMA", "CMF")

	for i, candle := range candleSticks {
		tbl.AddRow(time.UnixMilli(candle.OpenTime).UTC(), candle.Open, candle.Low, candle.High, candle.Close, (*ema)[i], (*cmf)[i])
	}

	tbl.Print()
}
