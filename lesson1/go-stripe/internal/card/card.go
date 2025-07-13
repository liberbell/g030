package card

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
