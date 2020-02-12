package sites

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/apiotrowski312/isOnline-sites-api/src/domain/sites"
	"github.com/apiotrowski312/isOnline-sites-api/src/services"
	"github.com/apiotrowski312/isOnline-sites-api/src/utils/http_utils"
	"github.com/apiotrowski312/isOnline-utils-go/oauth"
	"github.com/apiotrowski312/isOnline-utils-go/rest_errors"
	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	siteId, err := strconv.Atoi(strings.TrimSpace(vars["id"]))
	if err != nil {
		http_utils.RespondError(w, rest_errors.NewInternalServerError("Wrong index", err))
		return

	}

	site, getErr := services.SitesService.GetSite(int64(siteId))
	if getErr != nil {
		http_utils.RespondError(w, getErr)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, site)
}

func Post(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status())
		if a := json.NewEncoder(w).Encode(err); a != nil {
			fmt.Println("Error json: " + a.Error())
		}
		return
	}

	userId := oauth.GetCallerId(r)

	if userId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid access token")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var siteRequest sites.Site
	if err := json.Unmarshal(requestBody, &siteRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	siteRequest.UserId = userId

	result, createErr := services.SitesService.SaveSite(siteRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func GetUserSites(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status())
		if a := json.NewEncoder(w).Encode(err); a != nil {
			fmt.Println("Error json: " + a.Error())
		}
		return
	}

	userId := oauth.GetCallerId(r)

	if userId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid access token")
		http_utils.RespondError(w, respErr)
		return
	}

	result, createErr := services.SitesService.FindByOwner(userId)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}
