package nanorpc

import (
	"time"
)

type RequestActionType string

const (
	RequestActionTypeAccountBalance RequestActionType = "account_balance"
	RequestActionTypeAccountHistory RequestActionType = "account_history"
)


type RequesAction struct {
	Action  RequestActionType `json:"action"`
	Account Address           `json:"account"`
}

type RequesActionAccountHistory struct {
	Action  RequestActionType `json:"action"`
	Account Address           `json:"account"`
	Count   int64             `json:"count"`
	Raw     *bool             `json:"raw,omitempty"`
}

type AccountBalanceResponse struct {
	Balance Raw `json:"balance"`
	Pending Raw `json:"pending"`
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
