package model

// Need to run go get github.com/shopspring/decimal to install the decimal package
// Then go mod tidy to update the go.mod file
import "github.com/shopspring/decimal"

type OrderLines struct {
	OrderLineID int `json:"order_line_id" db:"order_line_id"`
	OrderID int `json:"order_id" db:"order_id"`
	ItemID int `json:"item_id" db:"item_id"`
	Quantity decimal.Decimal `json:"quantity" db:"quantity"`
	UnitPrice decimal.Decimal `json:"unit_price" db:"unit_price"`
	Remark *string `json:"remark" db:"remark"`
}