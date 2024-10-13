package blockchain

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type reqMessage struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

type roundTripperExt struct {
	c *http.Client
}

func (r roundTripperExt) RoundTrip(request *http.Request) (*http.Response, error) {
	rt := request.Clone(context.Background())
	body, _ := io.ReadAll(request.Body)
	// log.Printf("%s \n\n\n\n", body)
	_ = request.Body.Close()
	if len(body) > 0 {
		rt.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	var req reqMessage
	if err := json.Unmarshal(body, &req); err == nil {
		d2, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}
		rt.ContentLength = int64(len(d2))
		rt.Body = ioutil.NopCloser(bytes.NewBuffer(d2))
	}
	return r.c.Do(rt)
}

// NewEthereumClientFromFlag returns Ethereum client from flag variable, or error if occurs
func NewEthereumClient(ethereumNodeURL string) (*ethclient.Client, error) {
	cc := &http.Client{Transport: roundTripperExt{c: &http.Client{}}}
	r, err := rpc.DialHTTPWithClient(ethereumNodeURL, cc)
	if err != nil {
		return nil, err
	}
	client := ethclient.NewClient(r)
	return client, nil
}
