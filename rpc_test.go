package nanorpc_test

import (
	"os"
	"testing"

	"github.com/oiime/nanorpc"
	"github.com/stretchr/testify/assert"
)

func NewTestClient(t *testing.T) *nanorpc.Client {
	endpoint := os.Getenv("NANORPC_TEST_ENDPOINT")

	assert.NotEmpty(t, endpoint)
	return nanorpc.NewClient(endpoint)
}
