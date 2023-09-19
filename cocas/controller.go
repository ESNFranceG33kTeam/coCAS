package cocas

import (
	"encoding/json"
	"net/http"

	"github.com/ESNFranceG33kTeam/coCAS/logger"
	"gopkg.in/cas.v2"
)

var err error

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	if !cas.IsAuthenticated(r) {
		cas.RedirectToLogin(w, r)
		return
	}

	if r.URL.Path == "/logout" {
		cas.RedirectToLogout(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)

	if r.URL.Path == "/healthcheck" {
		err = json.NewEncoder(w).Encode(GetHealth())
	} else if r.URL.Path == "/reload" {
		ExternalLoadConf()
		logger.GetLogger().LogInfo("cocas", "Reloaded conf")
		err = json.NewEncoder(w).Encode(GalaxyUsers)
	} else if r.URL.Path == "/conf" {
		err = json.NewEncoder(w).Encode(GalaxyUsers)
	} else {
		err = json.NewEncoder(w).Encode(GetProfile(r))
	}

	if err != nil {
		logger.GetLogger().LogError("cocas", "problem with encoder.", err)
	} else {
		logger.GetLogger().LogInfo("cocas", "request GET : "+r.RequestURI)
	}
}
