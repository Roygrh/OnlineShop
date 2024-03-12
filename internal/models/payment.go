package models

type Payment struct {
	ID          string
	Amount      float64
	Currency    string
	CardNumber  string
	ExpiryMonth int
	ExpiryYear  int
	CVV         string
	Status      string
	MerchantID  string
	CustomerID  string
	Timestamp   time.Time
}
