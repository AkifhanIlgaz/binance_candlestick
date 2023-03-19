# binance_candlestick

A command-line program that retrieves candlestick information of a given pair and calculates EMA 20, CMF

# How it works ?

Clone this repository to your local machine.

```bash
git clone https://github.com/AkifhanIlgaz/binance_candlestick.git
```

Create .env file and add your Secret Key and API key

```bash
APIKEY = <YOUR_API_KEY>
SECRETKEY = <YOUR_SECRET_KEY>
```

Build the program

```bash
go build
```

Run the executable with flags

```bash
./binance_candlestick --pair <PAIR_NAME> --interval <INTERVAL> --limit <NUMBER_OF_CANDLES>
```

Example command to retrieve last candlestick information of BTCUSDT pair with 1h interval

```bash
./binance_candlestick --pair BTCUSDT --interval 1h --limit 1
```

![example](https://github.com/AkifhanIlgaz/binance_candlestick/blob/main/example.JPG)
