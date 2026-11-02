package model

import time "time"

type CartItems struct {
	CartItemID int `json:"cart_item_id" db:"cart_item_id"`
	UserID int `json:"user_id" db:"user_id"`
	ItemID int `json:"item_id" db:"item_id"`
	Quantity int `json:"quantity" db:"quantity"`
	Remark *string `json:"remark" db:"remark"`
	AddedAt time.Time `json:"added_at" db:"added_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}