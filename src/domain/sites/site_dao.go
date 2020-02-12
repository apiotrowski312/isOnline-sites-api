package sites

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/apiotrowski312/isOnline-sites-api/src/datasources/postgres/sites_db"
	"github.com/apiotrowski312/isOnline-utils-go/logger"
	"github.com/apiotrowski312/isOnline-utils-go/rest_errors"
)

const (
	queryGetSite         = `SELECT id, user_id, url, status, short_name, description, enabled FROM sites WHERE id=$1;`
	queryFindByOwnerSite = `SELECT id, user_id, url, status, short_name, description, enabled FROM sites WHERE user_id=$1;`
	queryInsertSite      = `INSERT INTO sites(user_id, url, status, short_name, description, enabled) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;`
)

func (site *Site) Get() rest_errors.RestErr {
	stmt, err := sites_db.Client.Prepare(queryGetSite)

	if err != nil {
		logger.Error("error when trying prepare get site statement", err)
		return rest_errors.NewInternalServerError("error when trying to get site", errors.New("database error"))
	}

	defer stmt.Close()
	result := stmt.QueryRow(site.Id)

	if getErr := result.Scan(&site.Id, &site.UserId, &site.Url, &site.Status, &site.ShortName, &site.Description, &site.Enabled); getErr != nil {
		logger.Error("error when trying get site", err)
		return rest_errors.NewInternalServerError(getErr.Error(), errors.New("database error"))
	}

	return nil
}

func (site *Site) Save() rest_errors.RestErr {
	saveErr := sites_db.Client.QueryRow(queryInsertSite, site.UserId, site.Url, site.Status, site.ShortName, site.Description, site.Enabled).Scan(&site.Id)

	if saveErr != nil {
		logger.Error("error when trying to save site statement", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save site", errors.New("database error"))
	}

	return nil
}

func (site *Site) Delete() rest_errors.RestErr {
	return nil
}

func (site *Site) FindByOwner() (Sites, rest_errors.RestErr) {
	stmt, err := sites_db.Client.Prepare(queryFindByOwnerSite)

	if err != nil {
		logger.Error("error when trying prepare findByOwner site statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get sites", errors.New("database error"))
	}

	defer stmt.Close()
	var rows *sql.Rows
	rows, err = stmt.Query(site.UserId)

	if err != nil {
		logger.Error("error when trying exec findByOwner site statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get sites", errors.New("database error"))
	}

	results := make([]Site, 0)
	logger.Info(fmt.Sprintf("%d", site.UserId))

	for rows.Next() {
		var site Site

		if scanErr := rows.Scan(&site.Id, &site.UserId, &site.Url, &site.Status, &site.ShortName, &site.Description, &site.Enabled); scanErr != nil {
			logger.Error("error when trying findByOwner site", scanErr)
			return nil, rest_errors.NewInternalServerError(scanErr.Error(), errors.New("database error"))
		}
		logger.Info(site.Url)

		results = append(results, site)
	}

	return results, nil
}
