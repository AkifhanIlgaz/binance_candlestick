package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	symbol, interval, limit := parseArgs()

	candleSticks := getCandleSticks(*symbol, *interval, *limit)

	ema, cmf := calculateEMAandCMF(candleSticks)

	printResult(candleSticks, &ema, &cmf)
}
