package main

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("hit the handler")
	if err != nil {
		if err != app.renderTemplate(w, r, "terminlal", nil); err != nil {
			app.infoLog.Println("hit the handler")}
		}
	}
}
