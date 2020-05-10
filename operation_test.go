package queryapp

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidOperation(t *testing.T) {
	var unknownOperation Operation = "Unknown"

	tt := []struct {
		op    Operation
		valid bool
	}{
		{BUY, true},
		{SELL, true},
		{unknownOperation, false},
	}

	for _, v := range tt {
		err := IsValidOperation(v.op)

		if v.valid {
			require.NoError(t, err)
			continue
		}

		require.Error(t, err)
		require.True(t, errors.Is(err, ErrUnknownOperation))
	}
}
