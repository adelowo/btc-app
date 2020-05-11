package coindesk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/adelowo/queryapp"
)

const (
	COINDESK_URL = "https://api.coindesk.com/v1/bpi/currentprice.json"
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) FetchPrice() (float64, error) {
	req, err := http.NewRequest(http.MethodGet, COINDESK_URL, bytes.NewBuffer(nil))
	if err != nil {
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("invalid HTTP error code.. %v", err)
	}

	var res response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, err
	}

	return res.Bpi.USD.RateFloat, nil
}

func New(cl *http.Client) (queryapp.Client, error) {
	if cl == nil {
		cl = &http.Client{
			Timeout: time.Second * 5, // Fail after 5 seconds
		}
	}

	return &Client{cl}, nil
}
