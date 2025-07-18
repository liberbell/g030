package cards

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

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
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
			msg = string(stripeErr.Code)
		}
		return nil, msg, err
	}
	return pi, "", nil
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""
	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your card was declined."
	case stripe.ErrorCodeExpiredCard:
		msg = "Your card is expired."
	case stripe.ErrorCodeInvalidCVC:
		msg = "Incorrect CVC code."
	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect ZIP code."
	case stripe.ErrorCodeAmountTooLarge:
		msg = "The Amount is too large to charge to your card."
	case stripe.ErrorCodeAmountTooSmall:
		msg = "The Amount is too small to charge to your card."
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Insufficient Balance."
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "Your Postal cord is invalid."
	default:
		msg = "Your card was declined."
	}
	return msg
}
