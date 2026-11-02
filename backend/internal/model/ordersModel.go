package model

import time "time"
// Need to run go get github.com/shopspring/decimal to install the decimal package
// Then go mod tidy to update the go.mod file
import "github.com/shopspring/decimal"

type OrderStatus string

const (
	OrderPlaced OrderStatus = "placed"
	OrderPaid OrderStatus = "paid"
	OrderPacked OrderStatus = "packed"
	OrderDelivering OrderStatus = "delivering"
	OrderPickUp OrderStatus = "ready_for_pickup"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "cancelled"
)

type FulfilmentType string

const (
	FulfilmentDelivery FulfilmentType = "delivery"
	FulfilmentPickUp FulfilmentType = "pickup"
)

type Orders struct {
	OrderID int `json:"order_id" db:"order_id"`
	UserID int `json:"user_id" db:"user_id"`
	OrderAt time.Time `json:"order_at" db:"order_at"`
	TotalCost decimal.Decimal `json:"total_cost" db:"total_cost"`
	Status OrderStatus `json:"status" db:"status"`
	FulfilmentType FulfilmentType `json:"fulfilment_type" db:"fulfilment_type"`
	DeliveryAddress *string `json:"delivery_address" db:"delivery_address"`
	DeliveryFee *decimal.Decimal `json:"delivery_fee" db:"delivery_fee"`
	PointsEarned int `json:"points_earned" db:"points_earned"`
	CompletedAt *time.Time `json:"completed_at" db:"completed_at"`
}