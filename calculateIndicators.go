package main

import (
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/cinar/indicator"
)

func calculateEMAandCMF(candleSticks []*binance.Kline) ([]float64, []float64) {
	var (
		highs    []float64
		lows     []float64
		closings []float64
		volumes  []int64
	)

	for _, candle := range candleSticks {
		highs = append(highs, parseFloat(candle.High))
		lows = append(lows, parseFloat(candle.Low))
		closings = append(closings, parseFloat(candle.Close))
		volumes = append(volumes, parseInt(candle.Volume))
	}

	ema := indicator.Ema(20, closings)
	cmf := indicator.ChaikinMoneyFlow(highs, lows, closings, volumes)

	return ema, cmf
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
