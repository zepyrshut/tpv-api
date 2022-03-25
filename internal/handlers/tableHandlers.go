package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/util"
)

var appTable *config.Application
var stTable *config.AppStatus

func NewTableHandler(a *config.Application, st *config.AppStatus) {
	appLounge = a
	stLounge = st
}

func GetTableFromLounge(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	loungeId, err := strconv.Atoi(params.ByName("id_salon"))
	if err != nil {
		appTable.Logger.Println(errors.New("invalid id parameter"))
		util.ErrorJSON(w, err)
		return
	}

	tables, err := appTable.DB.AllTablesFromSelectedLounge(loungeId)
	if err != nil {
		appTable.Logger.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, tables, "tables")
	if err != nil {
		appTable.Logger.Println(err)
		util.ErrorJSON(w, err)
		return
	}
}
