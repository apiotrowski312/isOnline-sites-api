package ticker

import (
	"fmt"
	"sync"
	"time"

	"github.com/apiotrowski312/isOnline-sites-api/src/services"
	"github.com/apiotrowski312/isOnline-sites-api/src/utils/checker"
	"github.com/apiotrowski312/isOnline-utils-go/logger"
)

func SetupTicker(tickers []int) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(tickers))
	for _, tick := range tickers {
		go addTicker(tick)
	}
	waitGroup.Wait()
}

func addTicker(tickerType int) {
	ticker := time.NewTicker(time.Duration(tickerType) * time.Second)

	for _ = range ticker.C {
		tickers, err := services.TickerService.FindByTickType(int64(tickerType))

		if err != nil {
			logger.Error("Ticker error", err)
		}

		for _, tick := range tickers {
			newStatus := checker.GetStatus(tick.Url)

			if newStatus != tick.Status {
				logger.Info(fmt.Sprintf("Status for %s has changed. %d to %d", tick.Url, newStatus, tick.Status))
				tick.Status = newStatus
				tick.Update()
			}
		}

	}
}
