package handlers

import (
	"net/http"

	"github.com/zepyrshut/tpv/internal/util"
)

func (m *Repository) GetAllLounges(w http.ResponseWriter, r *http.Request) {

	lounges, err := m.DB.AllLounges()
	if err != nil {
		m.App.ErrorLog.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, lounges, "lounges")

}
