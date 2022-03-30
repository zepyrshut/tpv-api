package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/zepyrshut/tpv/internal/util"
)

func (m *Repository) GetTableFromLounge(w http.ResponseWriter, r *http.Request) {

	loungeId, err := strconv.Atoi(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		m.App.ErrorLog.Println(errors.New("invalid id parameter"))
		util.ErrorJSON(w, err)
		return
	}

	tables, err := m.DB.AllTablesFromSelectedLounge(loungeId)
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, tables, "tables")
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}
}
