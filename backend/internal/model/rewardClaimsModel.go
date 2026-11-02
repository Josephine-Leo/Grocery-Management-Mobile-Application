package model

import time "time"

type RewardStatus string

const (
	RewardClaimed RewardStatus = "claimed"
	RewardUsed RewardStatus = "used"
	RewardExpired RewardStatus = "expired"
)

type RewardClaims struct {
	ClaimID int `json:"claim_id" db:"claim_id"`
	UserID int `json:"user_id" db:"user_id"`
	RewardID int `json:"reward_id" db:"reward_id"`
	Status RewardStatus `json:"status" db:"status"`
	ClaimedAt time.Time `json:"claimed_at" db:"claimed_at"`
	UsedAt *time.Time `json:"used_at" db:"used_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	OrderID *int `json:"order_id" db:"order_id"`
}