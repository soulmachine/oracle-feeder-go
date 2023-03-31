package internal

import (
	"fmt"
	"time"

	"github.com/terra-money/oracle-feeder-go/internal/fiat"
	internal_types "github.com/terra-money/oracle-feeder-go/internal/types"
	"github.com/terra-money/oracle-feeder-go/pkg/types"
	"golang.org/x/exp/maps"
)

type FiatProvider struct {
	priceBySymbol map[string]internal_types.PriceBySymbol
}

func NewFiatProvider(exchange string, symbols []string, interval int, timeout int, stopCh <-chan struct{}) (*FiatProvider, error) {
	fiatClient, err := fiat.GetFiatClient(exchange, symbols)
	if err != nil {
		return nil, err
	}

	provider := &FiatProvider{
		priceBySymbol: make(map[string]internal_types.PriceBySymbol),
	}

	go func() {
		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		prices, err := fiatClient.FetchAndParse(symbols, timeout)
		if err != nil {
			return
		}
		maps.Copy(provider.priceBySymbol, prices)
		for {
			select {
			case <-stopCh:
				ticker.Stop()
				return
			case <-ticker.C:
				prices, err := fiatClient.FetchAndParse(symbols, timeout)
				if err == nil {
					maps.Copy(provider.priceBySymbol, prices)
				}
			}
		}
	}()

	return provider, nil
}

func (p *FiatProvider) GetPrices() map[string]types.PriceByPair {
	result := make(map[string]types.PriceByPair)
	for _, price := range p.priceBySymbol {
		pair := fmt.Sprintf("%s/%s", price.Base, price.Quote)
		result[pair] = types.PriceByPair{
			Base:      price.Base,
			Quote:     price.Quote,
			Price:     price.Price,
			Timestamp: price.Timestamp,
		}
	}
	return result
}
