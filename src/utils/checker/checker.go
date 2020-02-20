package checker

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func GetStatus(url string) int64 {
	resp, err := http.Get(url)

	if err != nil {
		return -1
	}

	return int64(resp.StatusCode)
}

func Ticker() {
	ticker1s := time.NewTicker(1 * time.Second)
	ticker5s := time.NewTicker(5 * time.Second)
	ticker12s := time.NewTicker(12 * time.Second)
	done := make(chan bool)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		for {
			select {
			case <-done:
				waitGroup.Done()
				return
			case t := <-ticker1s.C:
				fmt.Println("ticker1s at", t)
			case t := <-ticker5s.C:
				fmt.Println("ticker5s at", t)
			case t := <-ticker12s.C:
				fmt.Println("ticker12s at", t)
			}
		}
	}()

	waitGroup.Wait()
}
