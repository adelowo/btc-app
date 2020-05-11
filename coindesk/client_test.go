// +build integration

package coindesk

import (
	"net/http"
	"testing"
	"time"

	"github.com/adelowo/queryapp"
	"github.com/stretchr/testify/require"
)

var _ queryapp.Client = (*Client)(nil)

func TestNew(t *testing.T) {

	// Test if nil is passed to the function
	c, err := New(nil)
	require.NoError(t, err)

	client := c.(*Client)

	require.NotNil(t, client.httpClient)

	// Test with a non client HTTP client
	const timeout = time.Minute * 10

	httpClient := &http.Client{
		Timeout: timeout,
	}

	c, err = New(httpClient)
	require.NoError(t, err)

	client = c.(*Client)

	require.NotNil(t, client.httpClient)

	require.Equal(t, timeout, client.httpClient.Timeout)
	require.Equal(t, httpClient, client.httpClient)
}
