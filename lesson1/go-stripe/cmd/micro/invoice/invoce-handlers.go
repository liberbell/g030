package main

import (
	"net/http"
	"time"
)

type Order struct {
	ID          int         `json: "id"`
	Quantity    int         `json: "quantity"`
	Amount      int         `json: "amount"`
	Product     string      `json: "product"`
	CreatedAt   time.Time   `json: "-"`
	UpdatedAt   time.Time   `json: "-"`
	Widget      Widget      `json: "widget"`
	Transaction Transaction `json: "transaction"`
	Customer    Customer    `json: "customer"`
}

func (app *application) CreateAndSendInvoce(w http.ResponseWriter, r *http.Request) {

}
