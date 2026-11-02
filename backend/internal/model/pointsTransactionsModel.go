package model

import time "time"

type PointsReason string

const (
	PointsEarned PointsReason = "earned"
	PointsRedeemed PointsReason = "redeemed"
	PointsAdjustment PointsReason = "adjustment"
)

type PointsTransactions struct {
	TxnID int `json:"txn_id" db:"txn_id"`
	UserID int `json:"user_id" db:"user_id"`
	OrderID *int `json:"order_id" db:"order_id"` // Can be redeem points so no order_id
	ClaimID *int `json:"claim_id" db:"claim_id"` // Can be redeem reward/earned points so no claim_id
	Change int `json:"change" db:"change"` // either +ve or -ve
	Reason PointsReason `json:"reason" db:"reason"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}