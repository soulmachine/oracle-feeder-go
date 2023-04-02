package exchangerate

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	internal_types "github.com/terra-money/oracle-feeder-go/internal/types"
)

const (
	exchange = "exchangerate"
	baseUrl  = "https://api.exchangerate.host/latest"
)

type ExchangeRateClient struct{}

func NewExchangeRateClient() *ExchangeRateClient {
	return &ExchangeRateClient{}
}

func (p *ExchangeRateClient) FetchAndParse(symbols []string, timeout int) (map[string]internal_types.PriceBySymbol, error) {
	var queryCurrencies []string
	for _, symbol := range symbols {
		items := strings.Split(symbol, "/")
		queryCurrencies = append(queryCurrencies, items[0])
	}
	url := fmt.Sprintf("%s?base=USD&symbols=%s", baseUrl, strings.Join(queryCurrencies, ","))
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	fmt.Printf("url: %s\n", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	jsonObj := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonObj)
	if err != nil {
		log.Printf("parse response error: %v\n", string(body))
		return nil, err
	}
	rates, ok := jsonObj["rates"].(map[string]interface{})
	if !ok {
		log.Printf("no rates: %v\n", string(body))
		return nil, err
	}
	prices := make(map[string]internal_types.PriceBySymbol)
	now := time.Now()
	quote := "USD"
	for k, v := range rates {
		// log.Printf("k: %s v: %v\n", k, v.(float64))
		base := k
		symbol := fmt.Sprintf("%s/USD", base)
		price := v.(float64)
		if price > 0 {
			prices[symbol] = internal_types.PriceBySymbol{
				Exchange:  exchange,
				Symbol:    symbol,
				Base:      base,
				Quote:     quote,
				Price:     1.0 / price,
				Timestamp: uint64(now.UnixMilli()),
			}
		}
	}
	return prices, nil
}
