// ./app/models/user.go

package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `db:"id" json:"id" validate:"required,uuid"`

	Username      string    `db:"username" json:"username" validate:"required,lte=64"`
	Password      string    `db:"password" json:"password" validate:"required"`
	AccountStatus bool      `db:"account_status" json:"account_status"`
	IsDeleted     bool      `db:"isDeleted" json:"isDeleted"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt     time.Time `db:"deleted_at" json:"deleted_at"`
}
