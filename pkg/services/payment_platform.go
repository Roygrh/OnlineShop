package services

import "payment/pkg/models"
import "payment/internal/models"

type PaymentPlatform struct {
	Name        string
	Description string
	SwiftCode   string
	Amount      float64
	db          *models.database
}

func NewPaymentPlatform(db *models.database) PaymentPlatform {
	return &PaymentPlatform{
		db: db,
	}
}

func (b *PaymentPlatform) ValidateCard(cardNum string) error {
	//...
}

func (b *PaymentPlatform) TransferFunds(amount float64) error {
	// ...
}

func (b *PaymentPlatform) SendPaymentDetails(payment models.Payment) error {
	//...
}
