package bitfinex

import (
	"fmt"
	"strings"

	"github.com/terra-money/oracle-feeder-go/configs"
)

func ParseSymbol(symbol string) (string, string, error) {
	if strings.HasPrefix(symbol, "t") {
		symbol = symbol[1:]
	}
	symbol = strings.ToUpper(symbol)
	for coin := range configs.FiatCoins {
		if strings.HasSuffix(symbol, coin) {
			base := strings.TrimSuffix(symbol, coin)
			return base, coin, nil
		}
	}
	for coin := range configs.StableCoins {
		if strings.HasSuffix(symbol, coin) {
			base := strings.TrimSuffix(symbol, coin)
			return base, coin, nil
		}
	}
	return "", "", fmt.Errorf("failed to parse Bitfinex %s", symbol)
}
