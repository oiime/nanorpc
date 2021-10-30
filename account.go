package nanorpc

func NewAccount(addr Address) *Account {
	return &Account{
		addr: addr,
	}
}

type Account struct {
	addr Address
}

func (a *Account) Balance(c *Client) (*AccountBalanceResponse, error) {
	rsp, err := c.Request(RequesAction{
		Action:  RequestActionTypeAccountBalance,
		Account: a.addr,
	})
	if err != nil {
		return nil, err
	}
	v := AccountBalanceResponse{}
	if err := rsp.Into(&v); err != nil {
		return nil, err
	}
	return &v, nil
}

func (a *Account) History(c *Client) (*AccountBalanceResponse, error) {
	rsp, err := c.Request(RequesActionAccountHistory{
		Action:  RequestActionTypeAccountHistory,
		Account: a.addr,
	})
	if err != nil {
		return nil, err
	}
	v := AccountBalanceResponse{}
	if err := rsp.Into(&v); err != nil {
		return nil, err
	}
	return &v, nil
}
