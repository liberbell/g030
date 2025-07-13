package card

import "github.com/stripe/stripe-go/v82"

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

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, error) {

}
