package server

import (
	"context"
	"pixare40/hello/proto/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetCurrency(ctx context.Context, rr *currency.CurrencyRequest) (*currency.CurrencyResponse, error) {
	c.log.Info("GetCurrency function was invoked", rr.GetCurrency(), rr.Currency)

	return &currency.CurrencyResponse{
		Currency: rr.Currency,
		Rate:     0.5,
	}, nil
}
