package services

import (
	"github.com/apiotrowski312/isOnline-sites-api/src/domain/sites"
	"github.com/apiotrowski312/isOnline-sites-api/src/utils/checker"
	"github.com/apiotrowski312/isOnline-utils-go/rest_errors"
)

var (
	SitesService SitesServiceInterface = &sitesService{}
)

type SitesServiceInterface interface {
	GetSite(int64) (*sites.Site, rest_errors.RestErr)
	SaveSite(sites.Site) (*sites.Site, rest_errors.RestErr)
	DeleteSite(int64) (*sites.Site, rest_errors.RestErr)
	FindByOwner(int64) (sites.Sites, rest_errors.RestErr)
}

type sitesService struct{}

func (s *sitesService) GetSite(siteId int64) (*sites.Site, rest_errors.RestErr) {
	results := &sites.Site{Id: siteId}

	if err := results.Get(); err != nil {
		return nil, err
	}

	return results, nil
}

func (s *sitesService) SaveSite(site sites.Site) (*sites.Site, rest_errors.RestErr) {

	status := checker.GetStatus(site.Url)

	site.Enabled = true
	site.Status = status

	if err := site.Save(); err != nil {
		return nil, err
	}

	return &site, nil
}

func (s *sitesService) DeleteSite(int64) (*sites.Site, rest_errors.RestErr) {
	return nil, nil
}

func (s *sitesService) FindByOwner(userId int64) (sites.Sites, rest_errors.RestErr) {
	results := &sites.Site{UserId: userId}

	sites, err := results.FindByOwner()

	if err != nil {
		return nil, err
	}

	return sites, nil
}
