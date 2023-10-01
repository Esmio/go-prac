// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package queries

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Kind string

const (
	KindExpenses Kind = "expenses"
	KindInCome   Kind = "in_come"
)

func (e *Kind) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Kind(s)
	case string:
		*e = Kind(s)
	default:
		return fmt.Errorf("unsupported scan type for Kind: %T", src)
	}
	return nil
}

type NullKind struct {
	Kind  Kind
	Valid bool // Valid is true if Kind is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullKind) Scan(value interface{}) error {
	if value == nil {
		ns.Kind, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Kind.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullKind) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Kind), nil
}

type Item struct {
	ID         int64     `json:"id"`
	UserID     int32     `json:"user_id"`
	Amount     int32     `json:"amount"`
	TagIds     []int32   `json:"tag_ids"`
	Kind       Kind      `json:"kind"`
	HappenedAt time.Time `json:"happened_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Tag struct {
	ID        int64      `json:"id"`
	UserID    int32      `json:"user_id"`
	Name      string     `json:"name"`
	Sign      string     `json:"sign"`
	Kind      Kind       `json:"kind"`
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type User struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ValidationCode struct {
	ID        int32      `json:"id"`
	Code      string     `json:"code"`
	Email     string     `json:"email"`
	UsedAt    *time.Time `json:"used_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
