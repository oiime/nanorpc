package nanorpc

import (
	"encoding/json"
	"math/big"
	"time"
)

type RequestActionType string

const (
	RequestActionTypeAccountBalance        RequestActionType = "account_balance"
	RequestActionTypeAccountBlockCount     RequestActionType = "account_block_count"
	RequestActionTypeAccountGet            RequestActionType = "account_get"
	RequestActionTypeAccountHistory        RequestActionType = "account_history"
	RequestActionTypeAccountInfo           RequestActionType = "account_info"
	RequestActionTypeAccountKey            RequestActionType = "account_key"
	RequestActionTypeAccountRepresentitive RequestActionType = "account_representative"
	RequestActionTypeAccountWeight         RequestActionType = "account_weight"
	RequestActionTypeAccountsBalances      RequestActionType = "accounts_balances"
	RequestActionTypeAccountsFrontiers     RequestActionType = "accounts_frontiers"
	RequestActionTypeAccountsPending       RequestActionType = "accounts_pending"
)

type RequestAction interface {
	RequestBody() ([]byte, error)
}

// RequestActionTypeAccountBalance
type AccountBalanceRequest struct {
	Account Address `json:"account"`
}

func (a AccountBalanceRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountBalanceRequest
		Action RequestActionType `json:"action"`
	}{
		AccountBalanceRequest: AccountBalanceRequest(a),
		Action:                RequestActionTypeAccountBalance,
	})
}

type AccountBalanceResponse struct {
	Balance Raw `json:"balance"`
	Pending Raw `json:"pending"`
}

// RequestActionTypeAccountBlockCount
type AccountBlockCountRequest struct {
	Account Address `json:"account"`
}

func (a AccountBlockCountRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountBlockCountRequest
		Action RequestActionType `json:"action"`
	}{
		AccountBlockCountRequest: AccountBlockCountRequest(a),
		Action:                   RequestActionTypeAccountBlockCount,
	})
}

type AccountBlockCountResponse struct {
	BlockCount int64 `json:"block_count"`
}

// RequestActionTypeAccountGet
type AccountGetRequest struct {
	Key string `json:"key"`
}

func (a AccountGetRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountGetRequest
		Action RequestActionType `json:"action"`
	}{
		AccountGetRequest: AccountGetRequest(a),
		Action:            RequestActionTypeAccountGet,
	})
}

type AccountGetResponse struct {
	Account Address `json:"account"`
}

// RequestActionTypeAccountHistory
type AccountHistoryRequest struct {
	Account Address `json:"account"`
	Count   int64   `json:"count"`
	Raw     *bool   `json:"raw,omitempty"`
}

func (a AccountHistoryRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountHistoryRequest
		Action RequestActionType `json:"action"`
	}{
		AccountHistoryRequest: AccountHistoryRequest(a),
		Action:                RequestActionTypeAccountHistory,
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
	Height         int64     `json:"height"`
	Hash           string    `json:"hash"`
}

// RequestActionTypeAccountInfo
type AccountInfoRequest struct {
	Account          Address `json:"account"`
	IncludeConfirmed bool    `json:"confirmed_balance"`
	Representative   bool    `json:"representative"`
	Weight           bool    `json:"weight"`
	Pending          bool    `json:"pending"`
}

func (a AccountInfoRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountInfoRequest
		Action RequestActionType `json:"action"`
	}{
		AccountInfoRequest: AccountInfoRequest(a),
		Action:             RequestActionTypeAccountInfo,
	})
}

type AccountInfoResponse struct {
	Frontier                   string    `json:"frontier"`
	OpenBlock                  string    `json:"open_block"`
	RepresentativeBlock        string    `json:"represntitive_block"`
	Balance                    Raw       `json:"balance"`
	ModifiedTimestamp          time.Time `json:"modified_timestamp"`
	AccountVersion             string    `json:"account_version"`
	ConfirmationHeight         *int64    `json:"confirmation_height"`
	ConfirmationHeightFrontier *string   `json:"confirmation_height_frontier"`
	ConfirmedBalance           *Raw      `json:"confirmed_balance"`
	ConfirmedHeight            *int64    `json:"confirmed_height"`
	ConfirmedFrontier          *string   `json:"confirmed_frontier"`
	Representative             *Address  `json:"representative"`
	Weight                     *big.Int  `json:"weight"`
	Pending                    *Raw      `json:"pending"`
}

// RequestActionTypeAccountKey
type AccountKeyRequest struct {
	Account Address `json:"account"`
}

func (a AccountKeyRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountKeyRequest
		Action RequestActionType `json:"action"`
	}{
		AccountKeyRequest: AccountKeyRequest(a),
		Action:            RequestActionTypeAccountKey,
	})
}

type AccountKeyResponse struct {
	Key string `json:"key"`
}

// RequestActionTypeAccountRepresentitive
type AccountRepresentitveRequest struct {
	Account Address `json:"account"`
}

func (a AccountRepresentitveRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountRepresentitveRequest
		Action RequestActionType `json:"action"`
	}{
		AccountRepresentitveRequest: AccountRepresentitveRequest(a),
		Action:                      RequestActionTypeAccountRepresentitive,
	})
}

type AccountRepresentitveResponse struct {
	Representative Address `json:"representative"`
}

// RequestActionTypeAccountWeight
type AccountWeightRequest struct {
	Account Address `json:"account"`
}

func (a AccountWeightRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountWeightRequest
		Action RequestActionType `json:"action"`
	}{
		AccountWeightRequest: AccountWeightRequest(a),
		Action:               RequestActionTypeAccountWeight,
	})
}

type AccountWeightResponse struct {
	Weight big.Int `json:"weight"`
}

// RequestActionTypeAccountsBalances
type AccountsBalancesRequest struct {
	Accounts []Address `json:"accounts"`
}

func (a AccountsBalancesRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountsBalancesRequest
		Action RequestActionType `json:"action"`
	}{
		AccountsBalancesRequest: AccountsBalancesRequest(a),
		Action:                  RequestActionTypeAccountsBalances,
	})
}

type AccountsBalancesResponse struct {
	Balances map[Address]AccountBalanceResponse `json:"balances"`
}

// RequestActionTypeAccountsFrontiers
type AccountsFrontiersRequest struct {
	Accounts []Address `json:"accounts"`
}

func (a AccountsFrontiersRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountsFrontiersRequest
		Action RequestActionType `json:"action"`
	}{
		AccountsFrontiersRequest: AccountsFrontiersRequest(a),
		Action:                   RequestActionTypeAccountsFrontiers,
	})
}

type AccountsFrontiersResponse struct {
	Frontiers map[Address]string `json:"frontiers"`
}

// RequestActionTypeAccountsPending
type AccountsPendingRequest struct {
	Accounts []Address `json:"accounts"`
	Count    int64     `json:"count"`
}

func (a AccountsPendingRequest) RequestBody() ([]byte, error) {
	return json.Marshal(struct {
		AccountsPendingRequest
		Action RequestActionType `json:"action"`
	}{
		AccountsPendingRequest: AccountsPendingRequest(a),
		Action:                 RequestActionTypeAccountsPending,
	})
}

// @TODO add request based type marshaling
type AccountsPendingResponse struct {
	Blocks map[Address][]string `json:"blocks"`
}
