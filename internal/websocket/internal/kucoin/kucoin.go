package kucoin

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/terra-money/oracle-feeder-go/internal/parser"
	"github.com/terra-money/oracle-feeder-go/internal/types"
)

const (
	kucoinToken  string = "2neAiuYvAU61ZDXANAGAsiL4-iAExhsBXZxftpOeh_55i3Ysy2q2LEsEWU64mdzUOPusi34M_wGoSf7iNyEWJ9UT_GlW6MKFx3WVnfdBDgj012qY6cO5GNiYB9J6i9GjsxUuhPw3BlrzazF6ghq4LzNy_YO9hlfLIw7mB4U2GIQ=.uhKVxeTcmIJteGNV6cDMWw=="
	websocketUrl string = "wss://ws-api-spot.kucoin.com/"
	exchangeName string = "kucoin"
)

type WebsocketClient struct{}

func NewWebsocketClient() *WebsocketClient {
	return &WebsocketClient{}
}

func (wc *WebsocketClient) ConnectAndSubscribe(symbols []string) (*websocket.Conn, error) {
	wsUrl := fmt.Sprintf("%s?token=%s&connectId=terra-price-server", websocketUrl, kucoinToken)
	conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		log.Println("#### line 31")
		return nil, err
	}

	for _, symbol := range symbols {
		command := generateCommand(symbol)
		if err := conn.WriteJSON(&command); err != nil {
			log.Println("#### line 38")
			return nil, err
		}

		resp := make(map[string]interface{})
		if err := conn.ReadJSON(&resp); err != nil {
			return nil, err
		}

		if status, ok := resp["type"].(string); ok && status == "error" {
			log.Println("#### line 46")
			bytes, _ := json.Marshal(resp)
			return nil, fmt.Errorf("%s", string(bytes))
		}
	}

	return conn, nil
}

// Candlestick websocket message.
//
// Message format: https://docs.kucoin.com/#klines
type RawCandlestickMsg struct {
	Type    string `json:"type"`
	Topic   string `json:"topic"`
	Subject string `json:"subject"`
	Data    struct {
		Symbol  string   `json:"symbol"`
		Candles []string `json:"candles"`
		TimeUs  uint64   `json:"time"`
	} `json:"data"`
}

// generateCommand generates the candlestick subscription command from specified symbols.
//
// API doc: https://docs.kucoin.com/#subscribe
//
// For example:
// {"sub":"market.btcusdt.kline.1min","id":"terra-price-server"}
func generateCommand(symbol string) map[string]interface{} {
	base, quote, err := parser.ParseSymbol("kucoin", symbol)
	if err != nil {
		log.Fatalf("symbol: %s invalid\n", symbol)
	}
	ex_symbol := fmt.Sprintf("%s-%s", base, quote)
	topic := fmt.Sprintf("/market/candles:%s_1min", ex_symbol)
	log.Printf("topic: %v\n", topic)
	return map[string]interface{}{
		"type":           "subscribe",
		"topic":          topic,
		"id":             "terra-price-server",
		"response":       false,
		"privateChannel": false,
	}
}

func convertSize(size string) float64 {
	result, err := strconv.ParseFloat(size, 8)
	if err != nil {
		return -1
	}
	return result
}

func (wc *WebsocketClient) ParseCandlestickMsg(rawMsg []byte) (*types.CandlestickMsg, error) {
	var msg RawCandlestickMsg
	json.Unmarshal(rawMsg, &msg)

	symbol := msg.Data.Symbol
	base, quote, err := parser.ParseSymbol(exchangeName, symbol)
	if err != nil {
		return nil, err
	}
	candles := msg.Data.Candles
	if len(candles) != 7 {
		return nil, fmt.Errorf("not candles %s", string(rawMsg))
	}
	// startTime := convertSize(candles[0])
	open, close, high, low := convertSize(candles[1]), convertSize(candles[2]), convertSize(candles[3]), convertSize(candles[4])
	vol, quoteVol := convertSize(candles[5]), convertSize(candles[6])
	vwap := 0.0
	if vol == 0.0 || quoteVol == 0.0 {
		vwap = (open + close) / 2.0
	} else {
		vwap = quoteVol / vol
	}

	return &types.CandlestickMsg{
		Exchange:  exchangeName,
		Symbol:    symbol,
		Base:      base,
		Quote:     quote,
		Timestamp: msg.Data.TimeUs / 1e3,
		Open:      open,
		High:      high,
		Low:       low,
		Close:     close,
		Volume:    vol,
		Vwap:      vwap,
	}, nil
}
