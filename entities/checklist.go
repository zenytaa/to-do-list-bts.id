package entities

import (
	"database/sql"
	"time"
)

type Checklist struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
