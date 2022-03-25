package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zepyrshut/tpv/internal/config"
)

var appHandler *config.Application
var stHanlder *config.AppStatus

func NewStatusHandler(a *config.Application, st *config.AppStatus) {
	appHandler = a
	stHanlder = st
}

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {

	currentStatus := config.AppStatus{
		Status:      "available",
		Environment: appHandler.Config.Env,
		Version:     stHanlder.Version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		appHandler.Logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}
