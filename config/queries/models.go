// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package queries

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int32
	Email     string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ValidationCode struct {
	ID        int32
	Code      string
	Email     string
	UsedAt    sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}
