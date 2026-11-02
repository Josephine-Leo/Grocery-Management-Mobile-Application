package model

import time "time"

type Rewards struct {
	RewardId int `json:reward_id" db:"reward_id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	PointsCost int `json:"points_cost" db:"points_cost"`
	ValidUntil time.Time `json:"valid_until" db:"valid_until"`
}

