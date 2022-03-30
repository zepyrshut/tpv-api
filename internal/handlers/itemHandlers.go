package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zepyrshut/tpv/internal/util"
)

func (m *Repository) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := m.DB.AllItems()
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, items, "items")
}

func (m *Repository) GetOneItem(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		m.App.ErrorLog.Println(errors.New("invalid id parameter"))
		util.ErrorJSON(w, err)
		return
	}

	item, err := m.DB.OneItem(id)
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, item, "item")
}