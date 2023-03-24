package websocket

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/terra-money/oracle-feeder-go/internal/types"
	"github.com/terra-money/oracle-feeder-go/internal/websocket/internal/binance"
	"github.com/terra-money/oracle-feeder-go/internal/websocket/internal/huobi"
	"github.com/terra-money/oracle-feeder-go/internal/websocket/internal/kucoin"
)

type websocketClient interface {
	ConnectAndSubscribe(symbols []string) (*websocket.Conn, error)
	ParseCandlestickMsg(msg []byte) (*types.CandlestickMsg, error)
}

func keepAlive(c *websocket.Conn, timeout time.Duration, setPong bool) {
	lastResponse := time.Now()
	if setPong {
		c.SetPongHandler(func(msg string) error {
			lastResponse = time.Now()
			return nil
		})
	}

	go func() {
		for {
			err := c.WriteMessage(websocket.PingMessage, []byte("keepalive"))
			if err != nil {
				log.Printf("Send ping error: %v", err)
			}
			time.Sleep(timeout / 2)
			if setPong {
				if time.Now().Sub(lastResponse) > timeout {
					log.Printf("Ping don't get response from %s", c.LocalAddr())
				}
			}
		}
	}()
}

// SubscribeCandlestick subscribes to the candlestick channel.
func SubscribeCandlestick(exchange string, symbols []string, stopCh <-chan struct{}) (<-chan *types.CandlestickMsg, error) {
	var client websocketClient
	switch strings.ToLower(exchange) {
	case "binance":
		client = binance.NewWebsocketClient()
	case "huobi":
		client = huobi.NewWebsocketClient()
	case "kucoin":
		client = kucoin.NewWebsocketClient()
	default:
		log.Printf("Exchange: %s not support\n", exchange)
		client = nil
	}

	conn, err := client.ConnectAndSubscribe(symbols)
	if err != nil {
		return nil, err
	}

	outCh := make(chan *types.CandlestickMsg)

	go func() {
		var rawMsg []byte
		for {
			select {
			case <-stopCh:
				conn.Close()
				close(outCh)
				return
			default:
				if conn != nil {
					_, rawMsg, err = conn.ReadMessage()
				}

				if err == nil {
					if exchange == "huobi" { // gzip decompress
						gzreader, _ := gzip.NewReader(bytes.NewReader(rawMsg))
						rawMsg, _ = io.ReadAll(gzreader)

						jsonObj := make(map[string]any)
						json.Unmarshal(rawMsg, &jsonObj)

						if status, ok := jsonObj["status"].(string); ok {
							if status != "ok" {
								log.Printf("%s", string(rawMsg))
							}
							break
						}
						if timestamp, ok := jsonObj["ping"]; ok {
							pongMsg := make(map[string]int64)
							pongMsg["pong"] = int64(timestamp.(float64))
							err = conn.WriteJSON(pongMsg)
							if err != nil {
								log.Printf("%v", err)
							}
							break
						}
					}

					candlestick, err := client.ParseCandlestickMsg(rawMsg)
					if err != nil {
						log.Printf("%v", err)
					} else {
						outCh <- candlestick
					}
				} else {
					log.Printf("line 59 %v %v", err, websocket.IsCloseError(err))
					if websocket.IsCloseError(err) {
						// reconnect automatically
						conn, err = client.ConnectAndSubscribe(symbols)
						if err != nil {
							log.Printf("%v", err)
							time.Sleep(3 * time.Second)
						}
					}
				}
			}
		}
	}()

	switch strings.ToLower(exchange) {
	case "kucoin":
		keepAlive(conn, time.Duration(time.Duration.Seconds(18)), false)
	}

	return outCh, nil
}
