package provider

import (
	"fmt"
	"strings"

	"github.com/terra-money/oracle-feeder-go/configs"
	"github.com/terra-money/oracle-feeder-go/internal/provider/internal"
	"github.com/terra-money/oracle-feeder-go/internal/provider/internal/coingecko"
	"github.com/terra-money/oracle-feeder-go/pkg/types"
)

// Provider represents a source of prices.
//
// For example, a cryptocurrency exchange can be a provider.
type Provider interface {
	GetPrices() map[string]types.PriceByPair
}

func NewProvider(exchange string, config *configs.ProviderConfig, stopCh <-chan struct{}) (Provider, error) {
	switch strings.ToLower(exchange) {
	case "binance", "huobi", "kucoin", "bitfinex", "kraken":
		return internal.NewWebsocketProvider(exchange, config.Symbols, stopCh)
	case "coingecko":
		return coingecko.NewCoingeckoProvider(config, stopCh)
	default:
		return nil, fmt.Errorf("unknown exchange %s", exchange)
	}
}
