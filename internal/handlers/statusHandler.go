package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zepyrshut/tpv/internal/config"
)

func (m *Repository) GetStatusHandler(w http.ResponseWriter, r *http.Request) {

	currentStatus := config.AppStatus{
		Status:      "OK",
		Environment: m.App.Status.Environment,
		Version:     m.App.Status.Version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		m.App.ErrorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}
