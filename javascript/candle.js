class CandleStick {
  constructor(candle) {
    [
      this.openTime,
      this.openPrice,
      this.highPrice,
      this.lowPrice,
      this.closePrice,
      this.volume,
      this.closeTime,
      this.assetVolume,
      this.trades,
      this.buyBaseVolume,
      this.buyAssetVolume,
      this.ignored,
    ] = candle.map((c) => +c);
  }
}

module.exports = CandleStick;
