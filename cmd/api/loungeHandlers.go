package main

import "net/http"

func (app *application) getAllLounges(w http.ResponseWriter, r *http.Request) {
	lounges, err := app.DB.AllLounges()
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w, err)
		return
	}

	app.WriteJSON(w, http.StatusOK, lounges, "lounges")

}
