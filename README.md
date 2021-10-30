# nanorpc

A collection of types and marshaling functionality for nanocurrency RPC protocol


    go get github.com/oiime/nanorpc


## Usage

```golang
import "github.com/oiime/nanorpc"

client := nanorpc.NewClient("https://proxy.powernode.cc/proxy")

addr := nanorpc.MustAddressFromString("nano_xyz")

// using helper types
act := nanorpc.NewAccount(addr)
rsp, err := act.Balance(client)
rsp, err := act.History(client)

// using types directly
rsp, err := client.Request(RequesActionAccountBalance{
    Action:  RequestActionTypeAccountBalance,
    Account: addr,
})

fmt.Printf("Nano: %s, Raw: %s\n", rsp.Balance.AsNanoString(), rsp.Balance.String())

rsp, err := client.Request(RequesActionAccountHistory{
    Action:  RequestActionTypeAccountHistory,
    Account: addr,
    Count: 1,
})

for _, entry range rsp.History {
    fmt.Printf("Hash: %s, Amount: %s\n", entry.Hash, entry.Amount.AsNanoString())
}

```

### Supported Actions

* RequestActionTypeAccountBalance "account_balance"
* RequestActionTypeAccountHistory "account_history"
