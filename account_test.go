package nanorpc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/oiime/nanorpc"
	"github.com/stretchr/testify/assert"
)

func TestAccountBalance(t *testing.T) {
	c := NewTestClient(t)
	act := nanorpc.NewAccount(nanorpc.MustAddressFromString(os.Getenv("NANORPC_TEST_ADDR")))
	rsp, err := act.Balance(c)
	assert.Nil(t, err)
	fmt.Println(rsp.Balance.AsNanoString(), rsp.Pending.AsNanoFloat())
}
