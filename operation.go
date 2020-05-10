package queryapp

import (
	"errors"
	"strings"
)

var ErrUnknownOperation = errors.New("invalid operation")

type Operation string

func (o Operation) String() string {
	return strings.ToLower(string(o))
}

const (
	BUY  Operation = "Buy"
	SELL Operation = "Sell"
)

func IsValidOperation(o Operation) error {
	switch o {
	case BUY, SELL:
		return nil
	default:
		return ErrUnknownOperation
	}
}
