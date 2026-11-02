package model

import time "time"
// Need to run go get github.com/shopspring/decimal to install the decimal package
// Then go mod tidy to update the go.mod file
import "github.com/shopspring/decimal"

type FridgeStatus string

const (
	InFridge FridgeStatus = "in_fridge"
	FridgeUsed FridgeStatus = "used"
	FridgeExpired FridgeStatus = "expired"
	FridgeDiscarded FridgeStatus = "discarded"
)

type FridgeEntries struct {
	EntryID int `json:"entry_id" db:"entry_id"`
	OrderLineID int `json:"order_line_id" db:"order_line_id"`
	QuantityRemaining decimal.Decimal `json:"quantity_remaining" db:"quantity_remaining"`
	ExpireAt *time.Time `json:"expire_at" db:"expire_at"`
	Status FridgeStatus `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}