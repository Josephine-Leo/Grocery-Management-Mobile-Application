package model

import time "time"
// Need to run go get github.com/shopspring/decimal to install the decimal package
// Then go mod tidy to update the go.mod file
import "github.com/shopspring/decimal"

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "pending"  // add nametag PaymentPending so can call model.PaymentPending to get the string "pending"
	PaymentSucceeded PaymentStatus = "succeeded"
	PaymentFailed PaymentStatus = "failed"
	PaymentRefunded PaymentStatus = "refunded"
)

// Add IsValid method to PaymentStatus type to check if the status is valid
// by calling (s PaymentStatus) creates a method similar to self in python, so can call model.PaymentStatus.IsValid() to check if the status is valid
func (s PaymentStatus) IsValid() bool {
    switch s {
    case PaymentPending, PaymentSucceeded, PaymentFailed, PaymentRefunded:
        return true
    }
    return false
}

type Payments struct {
	PaymentID   int       `json:"payment_id" db:"payment_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	OrderID	 int       `json:"order_id" db:"order_id"`
	Provider string    `json:"provider" db:"provider"`
	ProviderPaymentID string `json:"provider_payment_id" db:"provider_payment_id"`
	Amount decimal.Decimal `json:"amount" db:"amount"`
	PaidAt *time.Time `json:"paid_at" db:"paid_at"` 
	Status paymentStatus.PaymentStatus `json:"status" db:"status"`
}