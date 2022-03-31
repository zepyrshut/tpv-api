package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/zepyrshut/tpv/internal/util"
)

func (m *Repository) GetAllTypes(w http.ResponseWriter, r *http.Request) {
	types, err := m.DB.AllTypes()
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, types, "types")
}

func (m *Repository) GetOneType(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		m.App.ErrorLog.Println(errors.New("invalid id parameter"))
		util.ErrorJSON(w, err)
		return
	}

	typex, err := m.DB.OneType(id)
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, typex, "type")
}
