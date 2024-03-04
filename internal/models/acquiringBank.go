package models

import "payment/pkg/models"

type AcquiringBank struct {
	// properties
}

func (b *AcquiringBank) ValidateCard(cardNum string) error {
	//...
}

func (b *AcquiringBank) TransferFunds(amount float64) error {
	// ...
}

func (b *AcquiringBank) SendPaymentDetails(payment models.Payment) error {
	//...
}
