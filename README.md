# nanorpc

A collection of types and marshaling functionality for nanocurrency RPC protocol


    go get github.com/oiime/nanorpc


## Usage

There are two alternative= ways to generate calls, via helper types such as `Account` or passing an action request struct directly to the client

### using helper types
```golang
import "github.com/oiime/nanorpc"

client := nanorpc.NewClient("https://proxy.powernode.cc/proxy")
act := nanorpc.NewAccount(nanorpc.MustAddressFromString("nano_xyz"))
rsp, err := act.Balance(client)
fmt.Printf("Nano: %s, Raw: %s\n", rsp.Balance.AsNanoString(), rsp.Balance.String())


```

### using action types directly

```golang
import "github.com/oiime/nanorpc"

client := nanorpc.NewClient("https://proxy.powernode.cc/proxy")
addr := nanorpc.MustAddressFromString("nano_xyz")

rsp, err := client.Request(RequesActionAccountBalance{Account: addr})
fmt.Printf("Nano: %s, Raw: %s\n", rsp.Balance.AsNanoString(), rsp.Balance.String())

rsp, err := client.Request(RequesActionAccountHistory{Account: addr, Count: 1})
for _, entry range rsp.History {
    fmt.Printf("Hash: %s, Amount: %s\n", entry.Hash, entry.Amount.AsNanoString())
}

```

### Supported Actions

* RequestActionTypeAccountBalance "account_balance"
* RequestActionTypeAccountHistory "account_history"
