package fiat

import (
	"log"
	"strings"

	"github.com/terra-money/oracle-feeder-go/internal/fiat/internal/exchangerate"
	"github.com/terra-money/oracle-feeder-go/internal/fiat/internal/fer"
	"github.com/terra-money/oracle-feeder-go/internal/fiat/internal/frankfurter"
	internal_types "github.com/terra-money/oracle-feeder-go/internal/types"
)

type FiatClient interface {
	FetchAndParse(symbols []string, timeout int) (map[string]internal_types.PriceBySymbol, error)
}

func GetFiatClient(exchange string, symbols []string) (FiatClient, error) {
	var client FiatClient
	switch strings.ToLower(exchange) {
	case "exchangerate":
		client = exchangerate.NewFiatClient()
	case "frankfurter":
		client = frankfurter.NewFiatClient()
	case "fer":
		client = fer.NewFiatClient()
	default:
		log.Printf("fiat exchange: %s not support\n", exchange)
		client = nil
	}
	return client, nil
}
