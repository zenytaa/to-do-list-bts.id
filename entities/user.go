package entities

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Password  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
