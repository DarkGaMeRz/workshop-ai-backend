package main

import (
	"time"
)

type Transfer struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	FromUserID     uint      `json:"from_user_id" gorm:"not null"`
	ToUserID       uint      `json:"to_user_id" gorm:"not null"`
	Amount         int       `json:"amount" gorm:"not null;check:amount > 0"`
	Status         string    `json:"status" gorm:"not null;check:status IN ('pending','processing','completed','failed','cancelled','reversed')"`
	Note           *string   `json:"note" gorm:"type:text"`
	IdempotencyKey string    `json:"idempotency_key" gorm:"not null;unique"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CompletedAt    *time.Time `json:"completed_at"`
	FailReason     *string   `json:"fail_reason"`
}

type PointLedger struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	Change        int       `json:"change" gorm:"not null"`
	BalanceAfter  int       `json:"balance_after" gorm:"not null"`
	EventType     string    `json:"event_type" gorm:"not null;check:event_type IN ('transfer_out','transfer_in','adjust','earn','redeem')"`
	TransferID    *uint     `json:"transfer_id"`
	Reference     *string   `json:"reference"`
	Metadata      *string   `json:"metadata" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at"`
}