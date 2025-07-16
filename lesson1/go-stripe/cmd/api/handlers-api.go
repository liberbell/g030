package main

import (
	"encoding/json"
	"myapp/internal/cards"
	"net/http"
	"strconv"
)

type stripePayload struct {
	Currency string `json: "currency"`
	Amount   string `json: "amount"`
}

type jsonResponse struct {
	OK      bool   `json: "ok"`
	Message string `json: "message"`
	Content string `json: "content"`
	ID      int    `json: "id"`
}

func (app *application) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	var payload stripePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	amoount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	card := cards.Card{
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.key,
		Currency: payload.Currency,
	}

	j := jsonResponse{
		OK: true,
	}

	out, err := json.MarshalIndent(j, "", "   ")
	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
