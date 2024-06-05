package entities

import (
	"database/sql"
	"time"
)

type Cheklist struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
