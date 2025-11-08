package main

import (
	"net/http"
	"time"
)

type Order struct {
	ID        int       `json: "id"`
	Quantity  int       `json: "quantity"`
	Amount    int       `json: "amount"`
	Product   string    `json: "product"`
	CreatedAt time.Time `json: "created_at"`
}

func (app *application) CreateAndSendInvoice(w http.ResponseWriter, r *http.Request) {

}
