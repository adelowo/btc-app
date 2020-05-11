package queryapp

import (
	"errors"
	"strings"
)

var ErrUnknownOperation = errors.New("invalid BTC operation")

type Operation string

func (o Operation) String() string {
	return strings.ToLower(string(o))
}

const (
	UNKNOWNOperation           = ""
	BUY              Operation = "buy"
	SELL             Operation = "sell"
)

func IsValidOperation(o Operation) error {
	switch ParseOperation(o) {
	case BUY, SELL:
		return nil
	default:
		return ErrUnknownOperation
	}
}

func ParseOperation(o Operation) Operation {
	switch o.String() {
	case "buy":
		return BUY
	case "sell":
		return SELL
	default:
		return UNKNOWNOperation
	}
}
