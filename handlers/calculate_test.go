package handlers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	tt := []struct {
		btcPrice float64
		margin   float64
		plusOp   bool
		expected float64
	}{
		{100, 0.2, true, 100.2},
		{100.2, 0.2, true, 100.4004},
		{100, 0.2, false, 99.8},
		{100.2, 0.2, false, 99.9996},
	}

	for _, v := range tt {
		require.Equal(t, v.expected, calculate(v.btcPrice, v.margin, v.plusOp))
	}
}
