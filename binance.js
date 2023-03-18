const Binance = require("node-binance-api");
const CandleStick = require("./candle");
const {
  chaikinMoneyFlow,
  ema,
  bollingerBands,
  bollingerBandsWidth,
} = require("indicatorts");

require("dotenv").config();

const binance = new Binance().options({
  APIKEY: process.env.APIKEY,
  SECRETKEY: process.env.SECRETKEY,
  family: 4,
});

async function main() {
  let candlesticks = await binance.candlesticks("BTCUSDT", "1h");
  candlesticks = candlesticks.map((candle) => new CandleStick(candle));

  const EMA = ema(20, getPropertyFromCandles(candlesticks, "close"));
  const CMF = chaikinMoneyFlow(
    getPropertyFromCandles(candlesticks, "high"),
    getPropertyFromCandles(candlesticks, "low"),
    getPropertyFromCandles(candlesticks, "close"),
    getPropertyFromCandles(candlesticks, "volume")
  );
  const BollingerBands = bollingerBands(
    getPropertyFromCandles(candlesticks, "close")
  );
  const BollingerBandWidth = bollingerBandsWidth(BollingerBands);
}

function getPropertyFromCandles(candlesticks, property) {
  return candlesticks.map((candle) => candle[property]);
}

main();