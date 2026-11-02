package model

import time "time"
// Need to run go get github.com/shopspring/decimal to install the decimal package
// Then go mod tidy to update the go.mod file
import "github.com/shopspring/decimal"

type Items struct {
	ItemID int `json:"item_id" db:"item_id"`
	Name string `json:"name" db:"name"`
	SubCat string `json:"sub_cat" db:"sub_cat"`
	Price decimal.Decimal `json:"price" db:"price"`
	Discount string `json:"discount" db:"discount"`
	Rating string `json:"rating" db:"rating"`
	Title string `json:"title" db:"title"`
	Feature string `json:"feature" db:"feature"`
	Description string `json:"description" db:"description"`
	ItemImage string `json:"item_image" db:"item_image"` //url
}