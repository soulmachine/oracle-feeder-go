package coingecko

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/terra-money/oracle-feeder-go/internal/parser"
	internal_types "github.com/terra-money/oracle-feeder-go/internal/types"
)

const (
	baseUrl string = "https://api.coingecko.com/api/v3/simple/price"
)

type CoingeckoClient struct{}

func NewCoingeckoClient() *CoingeckoClient {
	return &CoingeckoClient{}
}

func (p *CoingeckoClient) FetchAndParse(symbols []string, timeout int) (map[string]internal_types.PriceBySymbol, error) {
	msg, err := fetchPrices(symbols)
	if err != nil {
		return nil, err
	}
	return parseJSON(msg), nil
}

func fetchPrices(symbols []string) (map[string]map[string]float64, error) {
	params := url.Values{}
	params.Add("vs_currencies", "usd")
	params.Add("precision", "18")
	params.Add("ids", strings.Join(symbols, ","))
	url := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	client := &http.Client{Timeout: time.Second * 15}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	jsonObj := make(map[string]map[string]float64)
	err = json.Unmarshal(body, &jsonObj)
	if err != nil {
		return nil, err
	}
	return jsonObj, nil
}

func parseJSON(msg map[string]map[string]float64) map[string]internal_types.PriceBySymbol {
	prices := make(map[string]internal_types.PriceBySymbol)
	now := time.Now()
	for symbol, value := range msg {
		base, quote, err := parser.ParseSymbol("coingecko", symbol)
		if err == nil {
			prices[symbol] = internal_types.PriceBySymbol{
				Exchange:  "coingecko",
				Symbol:    symbol,
				Base:      base,
				Quote:     quote,
				Price:     value["usd"],
				Timestamp: uint64(now.UnixMilli()),
			}
		} else {
			log.Printf("%v", err)
		}
	}
	return prices
}
