package main

import (
	"fmt"
	"net/http"
)

func (app *application) CreateAndSendInvoce(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Heelo %s", "world.")
}
