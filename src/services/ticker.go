package services

import (
	"github.com/apiotrowski312/isOnline-sites-api/src/domain/tickers"
)

var (
	TickerService TickerServiceInterface = &tickerService{}
)

type TickerServiceInterface interface {
	FindByTickType(int64) ([]tickers.Ticker, error)
}

type tickerService struct{}

func (s *tickerService) FindByTickType(tickType int64) ([]tickers.Ticker, error) {
	results := &tickers.Ticker{}

	tickers, err := results.FindByTickType(tickType)

	if err != nil {
		return nil, err
	}

	return tickers, nil
}
