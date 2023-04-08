package bittrex

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/terra-money/oracle-feeder-go/internal/parser"
	internal_types "github.com/terra-money/oracle-feeder-go/internal/types"
)

const (
	exchange = "bittrex"
	baseUrl  = "https://api.bittrex.com/v3"
)

type BittrexClient struct{}

func NewBittrexClient() *BittrexClient {
	return &BittrexClient{}
}

// Doc: https://bittrex.github.io/api/v3
// Candle api only support single currency pair, need fetch one by one.
func (p *BittrexClient) FetchAndParse(symbols []string, timeout int) (map[string]internal_types.PriceBySymbol, error) {
	prices := make(map[string]internal_types.PriceBySymbol)
	// log.Printf("start fetch")
	for _, symbol := range symbols {
		price, err := fetchSymbol(symbol, timeout)
		if err != nil {
			log.Printf("fetch symbol: %s error\n", symbol)
			continue
		}
		prices[symbol] = *price
	}
	// log.Printf("finish fetch")

	return prices, nil
}

func fetchSymbol(symbol string, timeout int) (*internal_types.PriceBySymbol, error) {
	url := fmt.Sprintf("%s/markets/%s/candles/trade/MINUTE_1/recent", baseUrl, symbol)
	// log.Printf("url: %s\n", url)
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	base, quote, err := parser.ParseSymbol(exchange, symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to parse symbol %s", symbol)
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var jsonObj []interface{}
	err = json.Unmarshal(body, &jsonObj)
	if err != nil || len(jsonObj) == 0 {
		return nil, fmt.Errorf("failed to unmarshal %s", string(body))
	}
	candle, ok := jsonObj[len(jsonObj)-1].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("no data: %s", string(body))
	}
	close, err := strconv.ParseFloat(candle["close"].(string), 64)
	if err != nil {
		return nil, err
	}
	open, err := strconv.ParseFloat(candle["open"].(string), 64)
	if err != nil {
		return nil, err
	}
	// high, err := strconv.ParseFloat(candle["high"].(string), 64)
	// if err != nil {
	// 	return nil, err
	// }
	// low, err := strconv.ParseFloat(candle["low"].(string), 64)
	// if err != nil {
	// 	return nil, err
	// }
	baseVolume, err := strconv.ParseFloat(candle["volume"].(string), 64)
	if err != nil {
		return nil, err
	}
	timeObj, err := time.Parse(time.RFC3339, candle["startsAt"].(string))
	if err != nil {
		return nil, err
	}
	timestamp := uint64(timeObj.UnixMilli())
	quoteVolume, err := strconv.ParseFloat(candle["quoteVolume"].(string), 64)
	if err != nil {
		return nil, err
	}
	vwap := 0.0
	if baseVolume == 0.0 || quoteVolume == 0.0 {
		vwap = (open + close) / 2.0
	} else {
		vwap = quoteVolume / baseVolume
	}
	price := &internal_types.PriceBySymbol{
		Exchange:  exchange,
		Symbol:    symbol,
		Base:      base,
		Quote:     quote,
		Price:     vwap,
		Timestamp: timestamp,
	}
	return price, nil
}
