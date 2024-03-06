package models

import "payment/internal/models"

type Customer struct {
    User models.User
    PaymentMethods []PaymentMethod
}