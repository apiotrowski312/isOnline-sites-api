package ping

import (
	"net/http"

	"github.com/apiotrowski312/isOnline-sites-api/src/utils/http_utils"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	http_utils.RespondJson(w, http.StatusOK, "pong")
}
