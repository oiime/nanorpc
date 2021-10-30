package nanorpc

import (
	"encoding/json"
	"time"
)

type RequestActionType string

const (
	RequestActionTypeAccountBalance RequestActionType = "account_balance"
	RequestActionTypeAccountHistory RequestActionType = "account_history"
)

type RequestAction interface {
	RequestBody() ([]byte, error)
}

// RequestActionTypeAccountBalance

type RequesActionBalance struct {
	Account Address `json:"account"`
}

func (a RequesActionBalance) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		RequesActionBalance
		Action RequestActionType `json:"action"`
	}{
		RequesActionBalance: RequesActionBalance(a),
		Action:              RequestActionTypeAccountBalance,
	})
}

type AccountBalanceResponse struct {
	Balance Raw `json:"balance"`
	Pending Raw `json:"pending"`
}

// RequestActionTypeAccountHistory

type RequesActionAccountHistory struct {
	Account Address `json:"account"`
	Count   int64   `json:"count"`
	Raw     *bool   `json:"raw,omitempty"`
}

func (a RequesActionAccountHistory) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		RequesActionAccountHistory
		Action RequestActionType `json:"action"`
	}{
		RequesActionAccountHistory: RequesActionAccountHistory(a),
		Action:                     RequestActionTypeAccountHistory,
	})
}

type AccountHistoryResponse struct {
	Account  Address               `json:"account"`
	History  []AccountHistoryEntry `json:"history"`
	Previous Block                 `json:"previous"`
}

type AccountHistoryEntry struct {
	Type           BlockType `json:"type"`
	Account        Address   `json:"account"`
	Amount         Raw       `json:"amount"`
	LocalTimestamp time.Time `json:"local_timestamp"`
	Height         Raw       `json:"height"`
	Hash           string    `json:"hash"`
}
