package nanorpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	Endpoint string
}

type Response struct {
	Body []byte
}

func (r *Response) Into(tgt interface{}) error {
	return json.Unmarshal(r.Body, tgt)
}

func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

func (c *Client) Request(act interface{}) (*Response, error) {
	payload, err := json.Marshal(act)
	if err != nil {
		panic(errors.Wrap(err, "NewRequestAction failed"))
	}

	rsp, err := http.Post(c.Endpoint, "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("received unexpected HTTP code %d", rsp.StatusCode)
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{
		Body: body,
	}, nil
}
