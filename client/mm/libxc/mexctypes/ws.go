// Package mexctypes provides type definitions for the MEXC exchange API.
package mexctypes

import (
	"encoding/json"
)

// WsRequest represents a WebSocket subscription request.
type WsRequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int64    `json:"id,omitempty"`
}

// WsSubResponse represents a WebSocket subscription response.
type WsSubResponse struct {
	ID     int64   `json:"id"`
	Result *bool   `json:"result"`
	Error  *string `json:"error,omitempty"`
}

// WsPBWrapper wraps protocol buffer encoded data in WebSocket messages.
type WsPBWrapper struct {
	Topic string          `json:"topic"`
	Data  []byte          `json:"data"`
	Time  int64           `json:"time,omitempty"`
	ID    json.RawMessage `json:"id,omitempty"`
}

// WsMessage represents a WebSocket message structure for MEXC.
type WsMessage struct {
	Type      string          `json:"type,omitempty"`
	Ping      int64           `json:"ping,omitempty"`
	Pong      int64           `json:"pong,omitempty"`
	Channel   string          `json:"channel,omitempty"` // Channel/topic name for subscription
	EventType string          `json:"e,omitempty"`       // Event type for user data
	Symbol    string          `json:"s,omitempty"`       // Symbol for market data
	Data      json.RawMessage `json:"data,omitempty"`    // Actual data payload
}

// WsPong is used to respond to ping messages to keep the websocket connection alive.
type WsPong struct {
	Pong int64 `json:"pong"`
}

// WsOrderUpdateData represents the order update data structure from the user data stream.
type WsOrderUpdateData struct {
	Symbol        string `json:"symbol"`              // Market symbol (e.g., BTCUSDT)
	OrderID       string `json:"orderId"`             // Exchange order ID
	ClientOrderID string `json:"clientOrderId"`       // Client-generated order ID
	Side          string `json:"side"`                // BUY or SELL
	Type          string `json:"type"`                // Order type (LIMIT, MARKET, etc.)
	Status        string `json:"status"`              // Order status (NEW, FILLED, CANCELED, etc.)
	Price         string `json:"price"`               // Order price
	Quantity      string `json:"quantity"`            // Original order quantity
	FilledQty     string `json:"executedQty"`         // Executed quantity
	FilledQuote   string `json:"cummulativeQuoteQty"` // Cumulative quote asset executed
	Time          int64  `json:"time"`                // Order creation time
	UpdateTime    int64  `json:"updateTime"`          // Last update time
}

// WsDealUpdateData represents the trade/deal update data structure from the user data stream.
type WsDealUpdateData struct {
	Symbol          string `json:"symbol"`          // Market symbol (e.g., BTCUSDT)
	DealID          string `json:"id"`              // Trade/Deal ID
	OrderID         string `json:"orderId"`         // Exchange order ID
	ClientOrderID   string `json:"clientOrderId"`   // Client-generated order ID
	Side            string `json:"side"`            // BUY or SELL
	Price           string `json:"price"`           // Trade price
	Quantity        string `json:"qty"`             // Trade quantity
	QuoteQty        string `json:"quoteQty"`        // Quote asset quantity
	Commission      string `json:"commission"`      // Commission amount
	CommissionAsset string `json:"commissionAsset"` // Commission asset
	Time            int64  `json:"time"`            // Trade execution time
}

// WsAccountUpdateData represents the account update data structure from the user data stream.
type WsAccountUpdateData struct {
	Asset     string `json:"a"` // Asset name
	Available string `json:"f"` // Free/Available balance
	Locked    string `json:"l"` // Locked/in-use balance
}

// WsDepthUpdateData represents an order book depth update from the market stream.
type WsDepthUpdateData struct {
	FirstUpdateID int64            `json:"F"` // First update ID in event
	FinalUpdateID int64            `json:"L"` // Final update ID in event
	Symbol        string           `json:"s"` // Symbol
	Bids          [][2]json.Number `json:"b"` // Bids [price, quantity]
	Asks          [][2]json.Number `json:"a"` // Asks [price, quantity]
}
