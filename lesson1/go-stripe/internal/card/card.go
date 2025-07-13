package card

import (
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionID  int
	Amount         int
	Currency       string
	LastFour       string
	BankReturnCode string
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, error, string) {
	stripe.Key = c.Secret

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}
	params.AddMetadata("key", "value")
	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = string(stripeError.Code)
		}
	}
}
