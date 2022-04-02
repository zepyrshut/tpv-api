package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/zepyrshut/tpv/internal/util"
)

func (m *Repository) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := m.DB.AllCategories()
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, categories, "categories")
}

func (m *Repository) GetOneCategory(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParamFromCtx(r.Context(), "id"))
	if err != nil {
		m.App.ErrorLog.Println(errors.New("invalid id parameter"))
		util.ErrorJSON(w, err)
		return
	}

	category, err := m.DB.OneCategory(id)
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, category, "category")
}
