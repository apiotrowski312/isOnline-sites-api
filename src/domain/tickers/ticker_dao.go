package tickers

import (
	"database/sql"
	"errors"

	"github.com/apiotrowski312/isOnline-sites-api/src/datasources/postgres/sites_db"
	"github.com/apiotrowski312/isOnline-utils-go/logger"
	"github.com/apiotrowski312/isOnline-utils-go/rest_errors"
)

const (
	queryFindByDurationType = `SELECT id, url, status FROM sites WHERE duration_type=$1 AND enabled=true;`
	queryUpdateStatus       = `UPDATE sites SET status=$1 WHERE id=$2;`
)

func (ticker *Ticker) Update() rest_errors.RestErr {
	stmt, err := sites_db.Client.Prepare(queryUpdateStatus)
	if err != nil {
		logger.Error("error when trying to prepare update ticker by id", err)
		return rest_errors.NewInternalServerError("error when trying to update ticker", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(ticker.Status, ticker.Id)

	if err != nil {
		logger.Error("error when trying to save site statement", err)
		return rest_errors.NewInternalServerError("error when trying to save site", errors.New("database error"))
	}

	return nil
}

func (ticker *Ticker) FindByTickType(tickType int64) ([]Ticker, rest_errors.RestErr) {
	stmt, err := sites_db.Client.Prepare(queryFindByDurationType)

	if err != nil {
		logger.Error("error when trying prepare queryFindByDurationType site statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get tickers", errors.New("database error"))
	}

	defer stmt.Close()
	var rows *sql.Rows
	rows, err = stmt.Query(tickType)

	if err != nil {
		logger.Error("error when trying exec queryFindByDurationType site statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get tickers", errors.New("database error"))
	}

	results := make([]Ticker, 0)

	for rows.Next() {
		var tick Ticker

		if scanErr := rows.Scan(&tick.Id, &tick.Url, &tick.Status); scanErr != nil {
			logger.Error("error when trying queryFindByDurationType site", scanErr)
			return nil, rest_errors.NewInternalServerError(scanErr.Error(), errors.New("database error"))
		}

		results = append(results, tick)
	}

	return results, nil
}
