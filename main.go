package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	symbol, interval, limit := parseArgs()

	candleSticks := getCandleSticks(*symbol, *interval, *limit)
	parseCandleData(candleSticks)
	ema, cmf, bbw := calculateEMA(), calculateCMF(), calculateBBW()

	printResult(candleSticks, &ema, &cmf, &bbw)
}
