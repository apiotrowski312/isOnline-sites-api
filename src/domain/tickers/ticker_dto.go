package tickers

type Ticker struct {
	Id     int64  `json:"id"`
	Url    string `json:"url"`
	Status int64  `json:"status"`
}

type Tickers []Ticker
