package entities

import (
	"database/sql"
	"time"
)

type Item struct {
	Id        int64
	ItemName  string
	Checklist Checklist
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
