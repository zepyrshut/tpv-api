package handlers

import (
	"net/http"

	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/util"
)

var appLounge *config.Application
var stLounge *config.AppStatus

func NewLoungeHandler(a *config.Application, st *config.AppStatus) {
	appLounge = a
	stLounge = st
}

func GetAllLounges(w http.ResponseWriter, r *http.Request) {

	lounges, err := appLounge.DB.AllLounges()
	if err != nil {
		appLounge.Logger.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, lounges, "lounges")

}
