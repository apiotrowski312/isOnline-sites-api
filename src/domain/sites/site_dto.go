package sites

import (
	"github.com/apiotrowski312/isOnline-sites-api/src/utils/checker"
	"github.com/apiotrowski312/isOnline-utils-go/rest_errors"
)

type Site struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	Url          string `json:"url"`
	Status       int64  `json:"status"`
	ShortName    string `json:"short_name"`
	Description  string `json:"description"`
	Enabled      bool   `json:"enabled"`
	DurationType int64  `json:"duration_type"`
}

type Sites []Site

func (site *Site) Validate() rest_errors.RestErr {
	status := checker.GetStatus(site.Url)

	if status == -1 {
		return rest_errors.NewBadRequestError("Wrong url")
	}

	site.Status = status

	return nil
}
