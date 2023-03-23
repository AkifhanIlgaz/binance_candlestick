package main

import (
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/cinar/indicator"
)

var (
	highs    []float64
	lows     []float64
	closings []float64
	volumes  []int64
)

func parseCandleData(candleSticks []*binance.Kline) {
	for _, candle := range candleSticks {
		highs = append(highs, parseFloat(candle.High))
		lows = append(lows, parseFloat(candle.Low))
		closings = append(closings, parseFloat(candle.Close))
		volumes = append(volumes, parseInt(candle.Volume))
	}
}

func calculateEMA() []float64 {
	return indicator.Ema(20, closings)
}

func calculateCMF() []float64 {
	return indicator.ChaikinMoneyFlow(highs, lows, closings, volumes)
}

func calculateBBW() []float64 {
	bandWidth, _ := indicator.BollingerBandWidth(indicator.BollingerBands(closings))
	return bandWidth
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
