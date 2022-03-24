package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getTableFromLounge(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	loungeId, err := strconv.Atoi(params.ByName("id_salon"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	tables, err := app.DB.AllTablesFromSelectedLounge(loungeId)
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}

	app.WriteJSON(w, http.StatusOK, tables, "tables")
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}
}
